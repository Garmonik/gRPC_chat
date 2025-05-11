package auth

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
)

func URLsWithoutVerification(cfg *config.Config, r *gin.RouterGroup, log *slog.Logger, dataBase *gorm.DB) {
	authHandlers := New(cfg, r, log, dataBase)

	r.GET("check/", authHandlers.healthCheck)
	r.POST("login/", authHandlers.loginHandler)
	r.POST("register/", authHandlers.registerHandler)

}

func URLs(cfg *config.Config, r *gin.RouterGroup, log *slog.Logger, dataBase *gorm.DB) {
	authHandlers := New(cfg, r, log, dataBase)

	r.GET("check/auth/", authHandlers.healthCheck)
	r.GET("logout/", authHandlers.logoutHandler)
}
