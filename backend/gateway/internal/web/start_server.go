package web

import (
	"context"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/web/handlers/auth"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/web/middleware"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"time"
)

func ConfigServer(cfg *config.Config, router *gin.Engine) *http.Server {
	return &http.Server{
		Addr:         cfg.HTTPServer.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}
}

func SetupRouter(log *slog.Logger, cfg *config.Config, dataBase *gorm.DB) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	public := router.Group("/api/v1")
	{
		auth.URLsWithoutVerification(cfg, public, log, dataBase)
	}

	private := router.Group("/api/v1")
	private.Use(middleware.AuthMiddleware(dataBase))
	{
		auth.URLs(cfg, public, log, dataBase)

	}

	log.Info("Starting API server",
		"address", cfg.HTTPServer.Address, "env", cfg.Environment)

	return router
}

func GracefulShutdown(server *http.Server, log *slog.Logger, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	log.Info("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		log.Error("Server forced to shutdown", "error", err)
	}

	log.Info("Server exited properly")

}
