package handlers

import (
	"net/http"

	"github.com/Smart-Door-Lock-IoT/api/features/configuration/dto/requests"
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/services"
	httpresponses "github.com/Smart-Door-Lock-IoT/api/pkg/http/responses"
	"github.com/gin-gonic/gin"
)

type Pin struct {
	service *services.Pin
}

func NewPin(
	service *services.Pin,
) *Pin {
	return &Pin{
		service: service,
	}
}

// @id 			ValidatePin
// @tags 		Configuration
// @param 		request body requests.ValidatePinRequest true "body"
// @success 	200 {object} responses.ValidatePinResponse
// @router 		/api/v1/configuration/pin [post]
func (h *Pin) ValidatePin(c *gin.Context) {
	var req requests.ValidatePinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.ValidatePin(req); err != nil {
		c.AbortWithStatusJSON(
			err.Code,
			httpresponses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

// @id 			ChangePin
// @tags 		Configuration
// @param 		request body requests.ChangePinRequest true "body"
// @success 	200 {object} responses.ChangePinResponse
// @router 		/api/v1/configuration/pin [put]
func (h *Pin) ChangePin(c *gin.Context) {
	var req requests.ChangePinRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if res, err := h.service.ChangePin(req); err != nil {
		c.AbortWithStatusJSON(
			err.Code,
			httpresponses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
