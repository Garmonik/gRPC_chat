package main

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/database"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/pkg/logger"
	_ "gorm.io/gorm"
	"log/slog"
)

func main() {
	// setup config
	cfg := config.MustLoad()

	// setup logger
	log := logger.SetupLogger(cfg.Environment)
	log.Info("starting auth services", slog.Any("Environment", cfg.Environment))

	db, err := database.New(log, cfg)
	if err != nil {
		log.Error("connection to the database failed", "error", err)
	}

}
