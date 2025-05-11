package main

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/database"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/pkg/logger_lib"

	"os"
)

func main() {
	cfg := config.MustLoad()
	log := logger_lib.SetupLogger(cfg.Environment)

	app := database.CreateMigratorApp(cfg, log)

	if err := app.Run(os.Args); err != nil {
		log.Error("migration failed", "error", err.Error())
	}
}
