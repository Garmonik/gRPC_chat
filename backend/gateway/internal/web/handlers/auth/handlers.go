package auth

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/grpc/auth_grpc"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/pkg/utils_lib"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"os"
	"strconv"
)

type Auth struct {
	cfg            *config.Config
	router         *gin.RouterGroup
	log            *slog.Logger
	DataBase       *gorm.DB
	grpcAuthClient *auth_grpc.GRPCAuthClient
}

func New(cfg *config.Config, r *gin.RouterGroup, log *slog.Logger, dataBase *gorm.DB) *Auth {
	grpcAuthClient, err := auth_grpc.New(log, cfg)
	if err != nil {
		log.Error("Failed to create grpc auth client", "error", err)
		os.Exit(2)
	}
	return &Auth{cfg: cfg, router: r, log: log, DataBase: dataBase, grpcAuthClient: grpcAuthClient}
}

func (a *Auth) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "available",
		"version": "1.0.0",
	})
}

func (a *Auth) loginHandler(ctx *gin.Context) {
	var data loginData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ipAddress := ctx.ClientIP()
	sessionUUID, err := a.grpcAuthClient.Login(ctx.Request.Context(), data.Email, data.Password, ipAddress)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	ctx.Header("Authorization", sessionUUID)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully logged in",
	})
}

func (a *Auth) registerHandler(ctx *gin.Context) {
	var data registerData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID, err := a.grpcAuthClient.Register(ctx.Request.Context(), data.Email, data.Password, data.Name)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "successfully registered in",
		"userID":  userID,
	})
}

func (a *Auth) logoutHandler(ctx *gin.Context) {
	sessionUUID := ctx.GetHeader("Authorization")
	user, err := utils_lib.GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
	}
	message, err := a.grpcAuthClient.Logout(ctx.Request.Context(), sessionUUID, strconv.Itoa(int(user.ID)))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (a *Auth) CloseSessionHandler(ctx *gin.Context) {
	var data closeSessionData
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := utils_lib.GetUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
	}
	message, err := a.grpcAuthClient.Logout(ctx.Request.Context(), data.SessionID, strconv.Itoa(int(user.ID)))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

func (a *Auth) sessionsList(ctx *gin.Context) {

}
