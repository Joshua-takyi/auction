package routes

import (
	"auction/internal/config"
	"auction/internal/container"
	"auction/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(cfg config.Config, container *container.Container) *gin.Engine {
	if cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	r := gin.New()
	r.Use(middleware.RequestID())
	r.Use(middleware.StructuredLogger(container.Logger))

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "health is working",
		})
	})
	// v1 := "/api/v0.0.0"

	return r
}
