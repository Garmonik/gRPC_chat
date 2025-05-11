package auth

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
)

type Auth struct {
	cfg      *config.Config
	router   *gin.RouterGroup
	log      *slog.Logger
	DataBase *gorm.DB
}

func New(cfg *config.Config, r *gin.RouterGroup, log *slog.Logger, dataBase *gorm.DB) *Auth {
	return &Auth{cfg: cfg, router: r, log: log, DataBase: dataBase}
}

func (a *Auth) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "available",
		"version": "1.0.0",
	})
}
