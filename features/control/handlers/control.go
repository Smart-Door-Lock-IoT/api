package handlers

import (
	"fmt"
	"net/http"

	"github.com/Smart-Door-Lock-IoT/api/features/control/dto/requests"
	"github.com/Smart-Door-Lock-IoT/api/features/control/services"
	httpresponses "github.com/Smart-Door-Lock-IoT/api/pkg/http/responses"
	"github.com/gin-gonic/gin"
)

type Control struct {
	service *services.Control
}

func NewControl(service *services.Control) *Control {
	return &Control{
		service: service,
	}
}

// @id 			TriggerOpenDoor
// @tags 		Control
// @success 	200 {object} responses.TriggerOpenDoorResponse
// @router 		/api/v1/control/open-door [post]
func (h *Control) TriggerOpenDoor(c *gin.Context) {
	if res, err := h.service.TriggerOpenDoor(); err != nil {
		c.AbortWithStatusJSON(
			err.Code, httpresponses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id 			TriggerBuzzerAlarm
// @tags 		Control
// @success 	200 {object} responses.TriggerBuzzerAlarmResponse
// @router 		/api/v1/control/buzzer-alarm [post]
func (h *Control) TriggerBuzzerAlarm(c *gin.Context) {
	if res, err := h.service.TriggerBuzzerAlarm(); err != nil {
		c.AbortWithStatusJSON(
			err.Code, httpresponses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id 			TriggerFingerprintMode
// @tags 		Control
// @param 		body body requests.TriggerFingerprintModeRequest true "body"
// @success 	200 {object} responses.TriggerFingerprintModeResponse
// @router 		/api/v1/control/fingerprint-mode [post]
func (h *Control) TriggerFingerprintMode(c *gin.Context) {
	var req requests.TriggerFingerprintModeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.TriggerFingerprintMode(req); err != nil {
		fmt.Println(err.Err)
		c.AbortWithStatusJSON(
			err.Code, httpresponses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id 			TriggerRFIDMode
// @tags 		Control
// @param 		body body requests.TriggerRFIDModeRequest true "body"
// @success 	200 {object} responses.TriggerRFIDModeResponse
// @router 		/api/v1/control/rfid-mode [post]
func (h *Control) TriggerRFIDMode(c *gin.Context) {
	var req requests.TriggerRFIDModeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.TriggerRFIDMode(req); err != nil {
		fmt.Println(err.Err)
		c.AbortWithStatusJSON(
			err.Code, httpresponses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
