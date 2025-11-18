package domains

import (
	"github.com/Smart-Door-Lock-IoT/api/features/control/models"
)

type Log struct {
	Id         uint   `json:"id"`
	DeviceName string `json:"device_name"`
	Status     bool   `json:"status"`
}

func (l *Log) ToModel() *models.Log {
	return &models.Log{
		DeviceName: l.DeviceName,
		Status:     l.Status,
	}
}

func FromLogModel(m *models.Log) *Log {
	return &Log{
		Id:         m.ID,
		DeviceName: m.DeviceName,
		Status:     m.Status,
	}
}
