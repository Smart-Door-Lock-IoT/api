package control

import (
	"github.com/Smart-Door-Lock-IoT/api/features/control/routes"
	"github.com/Smart-Door-Lock-IoT/api/injector"
	"github.com/gin-gonic/gin"
)

func RegisterControl(g *gin.RouterGroup) {
	handlers := injector.InitControlHandlers()

	routes.RegisterControl(g, handlers.Control)
}
