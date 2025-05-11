package main

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/database"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/pkg/logger_lib"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/web"
	_ "gorm.io/gorm"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// setup config
	cfg := config.MustLoad()

	// setup logger_lib
	log := logger_lib.SetupLogger(cfg.Environment)
	log.Info("starting auth services", slog.Any("Environment", cfg.Environment))

	db, err := database.New(log, cfg)
	if err != nil {
		log.Error("connection to the database failed", "error", err.Error())
	}

	router := web.SetupRouter(log, cfg, db.Db)
	server := web.ConfigServer(cfg, router)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Error("Server forced to shutdown", "error", err.Error())
			os.Exit(1)
		}
		log.Info("Server exited properly")

	}()

	// Graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signalStop := <-stop

	log.Info("stopping services", slog.String("signal", signalStop.String()))

	web.GracefulShutdown(server, log, 5*time.Second)
	err = db.Close()
	if err != nil {
		log.Error("close database failed", "error", err.Error())
		os.Exit(1)
	}
	log.Info("close database")
	log.Info("stop services")
}
