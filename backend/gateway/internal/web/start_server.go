package web

import (
	"context"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/web/handlers/auth"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/web/middleware"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/web/middleware/logger"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/web/middleware/recover"
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
	logger.SetupLoggingMiddleware(router, log)
	recover.SetupRecoveryMiddleware(router, log)

	setupPublicRoutes(router, log, cfg, dataBase)
	setupPrivateRoutes(router, log, cfg, dataBase)

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

func setupPublicRoutes(router *gin.Engine, log *slog.Logger, cfg *config.Config, db *gorm.DB) {
	public := router.Group("/api/v1")
	auth.URLsWithoutVerification(cfg, public, log, db)
}

func setupPrivateRoutes(router *gin.Engine, log *slog.Logger, cfg *config.Config, db *gorm.DB) {
	private := router.Group("/api/v1")
	private.Use(middleware.AuthMiddleware(db))
	auth.URLs(cfg, private, log, db)
}
