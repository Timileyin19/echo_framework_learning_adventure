package models

import (
	"time"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ProductID string
	Color string
	Quantity int64
	Timestamp time.Time
}