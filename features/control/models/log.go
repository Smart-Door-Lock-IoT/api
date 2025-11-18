package models

import "gorm.io/gorm"

type Log struct {
	gorm.Model

	DeviceName string `json:"device_name"`
	Status     bool   `json:"status"`
}
