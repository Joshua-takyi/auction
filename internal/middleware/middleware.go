package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := c.GetHeader("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}
		c.Set("request_id", requestID)
		c.Header("X-Request-ID", requestID)
		c.Next()
	}
}

func StructuredLogger(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next() // process the request

		latency := time.Since(start)
		statusCode := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		clientIP := c.ClientIP()
		requestID, _ := c.Get("request_id")

		if raw != "" {
			path += "?" + raw
		}

		// Decide log level based on status
		var logFunc func(msg string, args ...any)
		switch {
		case statusCode >= 500:
			logFunc = logger.Error
		case statusCode >= 400:
			logFunc = logger.Warn
		default:
			logFunc = logger.Info
		}

		logFunc("HTTP Request",
			slog.String("request_id", toString(requestID)),
			slog.String("method", method),
			slog.String("path", path),
			slog.Int("status", statusCode),
			slog.Duration("latency", latency),
			slog.String("client_ip", clientIP),
		)
	}
}

func toString(v any) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
