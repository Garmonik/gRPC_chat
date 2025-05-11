package recover

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func SetupRecoveryMiddleware(router *gin.Engine, log *slog.Logger) {
	router.Use(func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
				log.Error("Recovered from panic",
					slog.Any("error", err),
					slog.String("path", c.Request.URL.Path),
					slog.String("method", c.Request.Method),
					slog.String("ip", c.ClientIP()),
				)
			}
		}()
		c.Next()
	})
}
