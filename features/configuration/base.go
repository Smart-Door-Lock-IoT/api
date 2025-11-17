package configuration

import (
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/routes"
	"github.com/Smart-Door-Lock-IoT/api/injector"
	"github.com/gin-gonic/gin"
)

func RegisterConfiguration(g *gin.RouterGroup) {
	handlers := injector.InitConfigurationHandlers()

	routes.RegisterPin(g, handlers.Pin)
}
