package logger

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"time"
)

func SetupLoggingMiddleware(router *gin.Engine, log *slog.Logger) {
	router.Use(func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		errorMessage := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		attrs := []slog.Attr{
			slog.String("method", method),
			slog.String("path", path),
			slog.String("ip", clientIP),
			slog.Int("status", statusCode),
			slog.Duration("latency", latency),
			slog.String("user-agent", c.Request.UserAgent()),
		}

		if errorMessage != "" {
			attrs = append(attrs, slog.String("error", errorMessage))
		}

		log.LogAttrs(
			c.Request.Context(),
			getLogLevel(statusCode),
			"HTTP request",
			attrs...,
		)
	})
}

func getLogLevel(status int) slog.Level {
	switch {
	case status >= http.StatusInternalServerError:
		return slog.LevelError
	case status >= http.StatusBadRequest:
		return slog.LevelWarn
	default:
		return slog.LevelInfo
	}
}
