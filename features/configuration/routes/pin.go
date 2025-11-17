package routes

import (
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterPin(g *gin.RouterGroup, handler *handlers.Pin) {
	relativePath := "/configuration/pin"
	group := g.Group(relativePath)

	group.POST("", handler.ValidatePin)
	group.PUT("", handler.ChangePin)
}
