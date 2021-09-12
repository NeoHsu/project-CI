package health

import (
	"net/http"

	"github.com/NeoHsu/project-ci/internal/handlers"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func registerRoutes(handler *handlers.HttpHandler) {
	handler.Engine.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Health OK"})
	})
}

var Module = fx.Options(fx.Invoke(registerRoutes))
