package main

import (
	"github.com/Garmonik/gRPC_chat/backend/user/internal/app"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/pkg/logger_lib"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// setup config
	cfg := config.MustLoad()

	// setup logger_lib
	log := logger_lib.SetupLogger(cfg.Environment)
	log.Info("starting auth services", slog.Any("Environment", cfg.Environment))

	// setup application with gRPC
	application := app.New(log, cfg)

	go func() {
		err := application.GRPCServer.MustRun()
		if err != nil {
			log.Error("GRPC server failed to start", slog.Any("Error", err.Error()))
		}
	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signalStop := <-stop

	log.Info("stopping auth services", slog.String("signal", signalStop.String()))

	application.GRPCServer.Stop()

	log.Info("stop auth services")
}
