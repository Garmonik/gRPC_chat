package app

import (
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/app/grpcapp"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/general_settings/database"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/services/auth"
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

	authApp := auth.New(log, dataBase, cfg.SessionTTL, cfg)
	gRPCapp := grpcapp.New(log, authApp, cfg.GRPC.Port, cfg.GRPC.Host)
	return &App{GRPCServer: gRPCapp, DB: dataBase}
}
