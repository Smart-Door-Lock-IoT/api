package handlers

import "github.com/gin-gonic/gin"

type Control struct {
}

func NewControl() *Control {
	return &Control{}
}

func (h *Control) TriggerOpenDoor(c *gin.Context) {

}

func (h *Control) TriggerFingerprintMode(c *gin.Context) {

}

func (h *Control) TriggerRFIDMode(c *gin.Context) {

}
