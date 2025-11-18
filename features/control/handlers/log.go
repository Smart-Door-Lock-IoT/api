package handlers

import (
	"net/http"

	"github.com/Smart-Door-Lock-IoT/api/features/control/domains"
	"github.com/Smart-Door-Lock-IoT/api/features/control/dto/requests"
	"github.com/Smart-Door-Lock-IoT/api/features/control/services"
	httpresponses "github.com/Smart-Door-Lock-IoT/api/pkg/http/responses"
	"github.com/gin-gonic/gin"
)

type Log struct {
	service *services.Log
}

func NewLog(service *services.Log) *Log {
	return &Log{
		service: service,
	}
}

func (h *Log) CreateLog(req requests.CreateLogRequest) {
	h.service.CreateLog(
		domains.Log{
			DeviceName: req.DeviceName,
			Status:     req.Status,
		},
	)
}

func (h *Log) GetAllLogs(c *gin.Context) {
	if res, err := h.service.GetAllLogs(); err != nil {
		c.AbortWithStatusJSON(
			err.Code, httpresponses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}

func (h *Log) DeleteAllLogs(c *gin.Context) {
	if res, err := h.service.GetAllLogs(); err != nil {
		c.AbortWithStatusJSON(
			err.Code, httpresponses.Error{
				Message: err.Message,
			},
		)
	} else {
		c.JSON(http.StatusOK, res)
	}
}
