package auth

import (
	"context"
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/database"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/crypto_lib"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/interfase_lib"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/validate_lib"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/utils/auth_utils"
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
) (interfase_lib.SessionResponse, error) {
	const op = "Auth.Login"

	log := a.log.With(slog.String("op", op))
	log.Info("start Login service")
	defer log.Info("end Login service")

	if isValidEmail := validate_lib.ValidEmail(email); isValidEmail == false {
		log.Error("Email is invalid", slog.String("email", email))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false}, fmt.Errorf("email not valid")
	}
	user, err := auth_utils.GetUserByEmail(email, password, a.storage.Db)
	if err != nil {
		log.Error("User not found", slog.String("email", email), slog.String("error", err.Error()))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false}, err
	}
	isVerifyUser, err := crypto_lib.VerifyString(password, user.PasswordHash, a.cfg)
	if err != nil {
		log.Error("Error verifying password", slog.String("error", err.Error()))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false}, err
	}
	if isVerifyUser == false {
		log.Error("User is invalid", slog.String("email", email))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false}, fmt.Errorf("incorrect password")
	}
	sessionID, err := auth_utils.CreateNewSession(user.ID, a.storage.Db, a.cfg)
	if err != nil {
		log.Error("Error creating new session", slog.String("error", err.Error()))
		return interfase_lib.SessionResponse{SessionUUID: "", Success: false}, err
	}
	log.Info("session created", slog.String("sessionID", sessionID))
	return interfase_lib.SessionResponse{SessionUUID: sessionID, Success: true}, nil
}

func (a *Auth) RegisterNewUser(ctx context.Context, email string, password string, name string) (interfase_lib.RegisterResponse, error) {
	const op = "Auth.RegisterNewUser"
	log := a.log.With(slog.String("op", op))
	log.Info("start Register service")
	defer log.Info("end Register service")

	if isValidEmail := validate_lib.ValidEmail(email); isValidEmail == false {
		log.Error("Email is invalid", slog.String("email", email))
		return interfase_lib.RegisterResponse{UserId: 0, Code: interfase_lib.InvalidArgument}, fmt.Errorf("email not valid")
	}
	if isValidPassword := validate_lib.ValidPassword(password); isValidPassword == false {
		log.Error("Password is invalid", slog.String("password", password))
		return interfase_lib.RegisterResponse{UserId: 0, Code: interfase_lib.InvalidArgument}, fmt.Errorf("password not valid")
	}
	user, _ := auth_utils.GetUserByEmail(email, password, a.storage.Db)
	if user != nil {
		log.Error("User with this email already exists", slog.String("email", email))
		return interfase_lib.RegisterResponse{UserId: 0, Code: interfase_lib.AlreadyExists}, fmt.Errorf("user already exists")
	}
	user, err := auth_utils.CreateNewUser(email, password, name, a.storage.Db, a.cfg)

	return interfase_lib.RegisterResponse{}, nil
}

func (a *Auth) Logout(ctx context.Context, sessionUuid string) error {
	const op = "Auth.Logout"
	panic("implement me")
}
