package repositories

import (
	"github.com/Smart-Door-Lock-IoT/api/features/control/domains"
	"github.com/Smart-Door-Lock-IoT/api/features/control/models"
	"gorm.io/gorm"
)

type Log struct {
	db *gorm.DB
}

func NewLog(db *gorm.DB) *Log {
	return &Log{
		db: db,
	}
}

func (r *Log) Create(data domains.Log) (*domains.Log, error) {
	log := data.ToModel()
	if err := r.db.Create(&log).Error; err != nil {
		return nil, err
	} else {
		return domains.FromLogModel(log), nil
	}
}

func (r *Log) GetAll() (*[]domains.Log, error) {
	var logs []models.Log
	if err := r.db.Order("created_at desc").Find(&logs).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.Log, len(logs))
		for i, log := range logs {
			result[i] = *domains.FromLogModel(&log)
		}
		return &result, nil
	}
}

func (r *Log) GetAllLatest() (*[]domains.Log, error) {
	var logs []models.Log
	if err := r.db.Order("created_at desc").Limit(5).Find(&logs).Error; err != nil {
		return nil, err
	} else {
		result := make([]domains.Log, len(logs))
		for i, log := range logs {
			result[i] = *domains.FromLogModel(&log)
		}
		return &result, nil
	}
}

func (r *Log) DeleteAll() error {
	return r.db.Unscoped().Delete(&models.Log{}).Error
}
