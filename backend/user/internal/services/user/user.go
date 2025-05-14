package user

import (
	"context"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/database"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/database/models"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/pkg/interfase_lib"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/pkg/utils_lib/user_lib"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/pkg/validate_lib"
	"log/slog"
)

type ServiceUser struct {
	log     *slog.Logger
	storage *database.DataBase
	cfg     *config.Config
}

func New(
	log *slog.Logger,
	storage *database.DataBase,
	cfg *config.Config,
) *ServiceUser {
	return &ServiceUser{
		storage: storage,
		log:     log,
		cfg:     cfg,
	}
}

func (s *ServiceUser) GetMyUserInfo(
	ctx context.Context,
	userId int64,
) (*models.User, int) {
	const op = "ServiceUser.GetMyUserInfo"

	log := s.log.With(slog.String("op", op))
	log.Info("start GetMyUserInfo service")
	defer log.Info("end GetMyUserInfo service")

	user, err := user_lib.GetUserByID(ctx, uint64(userId), s.storage.Db)
	if err != nil {
		log.Error("User not found", slog.Int64("id", userId), slog.Any("error", err.Error()))
		return nil, interfase_lib.NotFound
	}

	log.Info("User found", slog.Int64("userID", userId))
	return user, interfase_lib.OK
}

func (s *ServiceUser) GetUserInfo(
	ctx context.Context,
	username string,
) (*models.User, int) {
	const op = "ServiceUser.GetUserInfo"

	log := s.log.With(slog.String("op", op))
	log.Info("start GetMyUserInfo service")
	defer log.Info("end GetMyUserInfo service")

	user, err := user_lib.GetUserByName(ctx, username, s.storage.Db)
	if err != nil {
		log.Error("User not found", slog.String("username", username), slog.Any("error", err.Error()))
		return nil, interfase_lib.NotFound
	}

	log.Info("User found", slog.String("username", username))
	return user, interfase_lib.OK
}

func (s *ServiceUser) UserUpdate(
	ctx context.Context,
	userId int64,
	email string,
	bio string,
) (int, string) {
	const op = "ServiceUser.UserUpdate"

	log := s.log.With(slog.String("op", op))
	log.Info("start UserUpdate service")
	defer log.Info("end UserUpdate service")

	checkValidEmail := validate_lib.ValidEmail(email)
	if checkValidEmail == false {
		log.Error("email is invalid", slog.String("email", email))
		return interfase_lib.InvalidArgument, "invalid email"
	}
	checkValidEmail = validate_lib.IsSafeInput(email)
	if checkValidEmail == false {
		log.Error("email is invalid", slog.String("email", email))
		return interfase_lib.InvalidArgument, "email is invalid"
	}
	checkValidBio := validate_lib.IsSafeInput(bio)
	if checkValidBio == false {
		log.Error("bio is invalid", slog.String("email", email))
		return interfase_lib.InvalidArgument, "bio is invalid"
	}

	typeError, err := user_lib.UpdateUser(ctx, userId, email, bio, s.storage.Db)
	switch typeError {
	case 0:
		return interfase_lib.OK, "change luck"
	case 1:
		log.Error("error with database", slog.String("error", err.Error()))
		return interfase_lib.Internal, "error with database"
	case 2:
		log.Error("user not found")
		return interfase_lib.NotFound, "user not found"
	default:
		return interfase_lib.Internal, "error with server"
	}
}
