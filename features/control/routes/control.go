package routes

import (
	"github.com/Smart-Door-Lock-IoT/api/features/control/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterControl(g *gin.RouterGroup, handler *handlers.Control) {
	relativePath := "/control"

	group := g.Group(relativePath)
	group.POST("/open-door", handler.TriggerOpenDoor)
	group.POST("/fingerprint-mode", handler.TriggerFingerprintMode)
	group.POST("/rfid-mode", handler.TriggerRFIDMode)
}
