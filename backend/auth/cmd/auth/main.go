package main

import (
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/config"
	"github.com/Garmonik/gRPC_chat/backend/auth/pkg/logger"
)

func main() {
	// setup config
	cfg := config.MustLoad()

	// setup base_logger
	log := logger.SetupLogger(cfg.Environment)
	log.Info("starting auth service")
	fmt.Println(cfg)
}
