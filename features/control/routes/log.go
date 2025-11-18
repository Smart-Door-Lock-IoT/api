package routes

import (
	"github.com/Smart-Door-Lock-IoT/api/features/control/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterLog(g *gin.RouterGroup, handler *handlers.Log) {
	relativePath := "/control/logs"

	group := g.Group(relativePath)
	group.GET("", handler.GetAllLogs)
	group.DELETE("", handler.DeleteAllLogs)
}
