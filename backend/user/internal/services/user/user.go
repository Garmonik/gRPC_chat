package user

import (
	"context"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/database"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/database/models"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/pkg/interfase_lib"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/pkg/utils_lib/auth_utils"
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

	user, err := auth_utils.GetUserByID(ctx, uint64(userId), s.storage.Db)
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
	const op = "ServiceUser.GetMyUserInfo"

	log := s.log.With(slog.String("op", op))
	log.Info("start GetMyUserInfo service")
	defer log.Info("end GetMyUserInfo service")

	user, err := auth_utils.GetUserByName(ctx, username, s.storage.Db)
	if err != nil {
		log.Error("User not found", slog.String("username", username), slog.Any("error", err.Error()))
		return nil, interfase_lib.NotFound
	}

	log.Info("User found", slog.String("username", username))
	return user, interfase_lib.OK
}
