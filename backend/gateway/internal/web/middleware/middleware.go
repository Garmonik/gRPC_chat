package middleware

import (
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/database/accessing_database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func AuthMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader("Authorization")
		fmt.Println("sessionID:", sessionID)
		if sessionID == "" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		sessionUUID, err := uuid.Parse(sessionID)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}
		session, err := accessing_database.GeSessionBySessionID(sessionUUID, db)
		if err != nil {
			c.Header("Authorization", "")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}
		if session.IsExpired() {
			session.Close()
			_ = db.Save(&session).Error

			c.Header("Authorization", "")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		if session.IsClosed {
			c.Header("Authorization", "")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
			return
		}

		c.Set("user", &session.User)
		c.Next()
	}
}
