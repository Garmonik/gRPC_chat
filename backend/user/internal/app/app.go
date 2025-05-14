package app

import (
	"github.com/Garmonik/gRPC_chat/backend/user/internal/app/grpcapp"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/database"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/services/user"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
	DB         *database.DataBase
}

func New(
	log *slog.Logger,
	cfg *config.Config) *App {

	dataBase, err := database.New(log, cfg)
	if err != nil {
		log.Error("Failed to connect to database", "error", err.Error())
	}

	userApp := user.New(log, dataBase, cfg)
	gRPCApp := grpcapp.New(log, userApp, cfg.GRPC.Port, cfg.GRPC.Host)
	return &App{GRPCServer: gRPCApp, DB: dataBase}
}
