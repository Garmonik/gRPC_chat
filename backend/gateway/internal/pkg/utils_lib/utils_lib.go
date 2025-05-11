package utils_lib

import (
	"errors"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/database/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUser(ctx *gin.Context) (*models.User, error) {
	userValue, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "user not found in context"})
		return nil, errors.New("user not found in context")
	}
	user, ok := userValue.(*models.User)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid user type in context"})
		return nil, errors.New("invalid user type in context")
	}
	return user, nil
}
