package main

import (
	"github.com/Garmonik/gRPC_chat/backend/auth/pkg/logger"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/database"

	"os"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Environment)

	app := database.CreateMigratorApp(cfg, log)

	if err := app.Run(os.Args); err != nil {
		log.Error("migration failed", "error", err.Error())
	}
}
