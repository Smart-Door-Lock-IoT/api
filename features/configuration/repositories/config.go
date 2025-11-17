package repositories

import (
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/domains"
	"github.com/Smart-Door-Lock-IoT/api/features/configuration/models"
	"gorm.io/gorm"
)

type Config struct {
	db *gorm.DB
}

func NewConfig(db *gorm.DB) *Config {
	return &Config{
		db: db,
	}
}

func (r *Config) Upsert(data domains.Config) (*domains.Config, error) {
	var config models.Config
	if err := r.db.Where(
		&domains.Config{
			Key: data.Key,
		},
	).Assign(
		&domains.Config{
			Value: data.Value,
		},
	).FirstOrCreate(&config).Error; err != nil {
		return nil, err
	} else {
		return domains.FromConfigModel(&config), nil
	}
}

func (r *Config) GetByKey(key string) (*domains.Config, error) {
	var config models.Config
	if err := r.db.Where("key = ?", key).First(&config).Error; err != nil {
		return nil, err
	} else {
		return domains.FromConfigModel(&config), nil
	}
}
