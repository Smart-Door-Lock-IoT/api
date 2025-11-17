package domains

import (
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/models"
)

type Config struct {
	Id    uint   `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (c *Config) ToModel() *models.Config {
	return &models.Config{
		Key:   c.Key,
		Value: c.Value,
	}
}

func FromConfigModel(m *models.Config) *Config {
	return &Config{
		Id:    m.ID,
		Key:   m.Key,
		Value: m.Value,
	}
}
