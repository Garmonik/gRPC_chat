package auth

import (
	"context"
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/database"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/crypto_lib"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/interfase_lib"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/utils_lib/auth_utils"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/validate_lib"
	"log/slog"
	"time"
)

type Auth struct {
	log        *slog.Logger
	storage    *database.DataBase
	sessionTTL time.Duration
	cfg        *config.Config
}

func New(
	log *slog.Logger,
	storage *database.DataBase,
	sessionTTL time.Duration,
	cfg *config.Config,
) *Auth {
	return &Auth{
		storage:    storage,
		log:        log,
		sessionTTL: sessionTTL,
		cfg:        cfg,
	}
}

func (a *Auth) Login(
	ctx context.Context,
	email string,
	password string,
	ipAddress string,
) (interfase_lib.SessionResponse, error) {
	const op = "Auth.Login"

	log := a.log.With(slog.String("op", op))
	log.Info("start Login service")
	defer log.Info("end Login service")

	if isValidEmail := validate_lib.ValidEmail(email); isValidEmail == false {
		log.Error("Email is invalid", slog.String("email", email))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false, Code: interfase_lib.InvalidArgument},
			fmt.Errorf("email not valid")
	}
	user, err := auth_utils.GetUserByEmail(email, a.storage.Db)
	if err != nil {
		log.Error("User not found", slog.String("email", email), slog.Any("error", err.Error()))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false, Code: interfase_lib.NotFound}, err
	}
	isVerifyUser, err := crypto_lib.VerifyString(password, user.PasswordHash, a.cfg)
	if err != nil {
		log.Error("Error verifying password", slog.String("error", err.Error()))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false, Code: interfase_lib.InvalidArgument}, err
	}
	if isVerifyUser == false {
		log.Error("User is invalid", slog.String("email", email))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false, Code: interfase_lib.InvalidArgument},
			fmt.Errorf("incorrect password")
	}
	sessionID, err := auth_utils.CreateNewSession(user.ID, ipAddress, a.storage.Db, a.cfg)
	if err != nil {
		log.Error("Error creating new session", slog.String("error", err.Error()))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false, Code: interfase_lib.Internal}, err
	}
	log.Info("session created", slog.String("sessionID", sessionID))
	return interfase_lib.SessionResponse{SessionUUID: sessionID, Success: true, Code: interfase_lib.OK}, nil
}

func (a *Auth) RegisterNewUser(
	ctx context.Context,
	email string,
	password string,
	name string,
) (interfase_lib.RegisterResponse, error) {
	const op = "Auth.RegisterNewUser"

	log := a.log.With(slog.String("op", op))
	log.Info("start Register service")
	defer log.Info("end Register service")

	if isValidEmail := validate_lib.ValidEmail(email); isValidEmail == false {
		log.Error("Email is invalid", slog.String("email", email))
		return interfase_lib.RegisterResponse{UserId: 0, Code: interfase_lib.InvalidArgument},
			fmt.Errorf("email not valid")
	}
	if isValidPassword := validate_lib.ValidPassword(password); isValidPassword == false {
		log.Error("Password is invalid", slog.String("password", password))
		return interfase_lib.RegisterResponse{UserId: 0, Code: interfase_lib.InvalidArgument},
			fmt.Errorf("password not valid")
	}
	user, _ := auth_utils.GetUserByEmail(email, a.storage.Db)
	if user != nil {
		log.Error("User with this email already exists", slog.String("email", email))
		return interfase_lib.RegisterResponse{UserId: 0, Code: interfase_lib.AlreadyExists},
			fmt.Errorf("user already exists")
	}
	user, _ = auth_utils.GetUserByName(name, a.storage.Db)
	if user != nil {
		log.Error("User with this email already exists", slog.String("name", name))
		return interfase_lib.RegisterResponse{UserId: 0, Code: interfase_lib.AlreadyExists},
			fmt.Errorf("user already exists")
	}
	user, err := auth_utils.CreateNewUser(email, password, name, a.storage.Db, a.cfg)
	if err != nil {
		log.Error("Error creating new user", slog.String("error", err.Error()))
		return interfase_lib.RegisterResponse{UserId: 0, Code: interfase_lib.Internal}, err
	}
	log.Info("Created new user", slog.String("email", email), slog.Uint64("id", uint64(user.ID)))
	return interfase_lib.RegisterResponse{UserId: int64(user.ID), Code: interfase_lib.OK}, nil
}

func (a *Auth) Logout(
	ctx context.Context,
	sessionUuid string,
	userId string) int {
	const op = "Auth.Logout"

	log := a.log.With(slog.String("op", op))
	log.Info("start Logout service")
	defer log.Info("end Logout service")

	userID, err := validate_lib.ConversionStringToUint(userId)
	if err != nil {
		log.Error("Error in ConversionStringToUint",
			slog.Any("error", err.Error()),
			slog.Any("userID", userID))
		return interfase_lib.InvalidArgument
	}
	user, err := auth_utils.GetUserByID(userID, a.storage.Db)
	if user == nil {
		log.Error("User not found",
			slog.Any("userID", userID))
		return interfase_lib.NotFound
	}
	session, err := auth_utils.CheckSessionID(sessionUuid, userID, a.storage.Db)
	if session == nil {
		log.Error("User does not have access to this session",
			slog.Any("error", err),
			slog.Any("userID", userID))
		return interfase_lib.PermissionDenied
	}
	err = auth_utils.CloseSession(sessionUuid, userID, a.storage.Db)
	if err != nil {
		log.Error("Error closing session",
			slog.Any("error", err.Error()),
			slog.Any("userID", userID))
		return interfase_lib.Internal
	}
	return interfase_lib.OK
}

func (a *Auth) SessionList(
	ctx context.Context,
	userId string) int {
	const op = "Auth.SessionList"
	panic("implement me")
}
