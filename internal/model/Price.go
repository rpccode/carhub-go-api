package model

import (
	"time"
)

type Price struct {
	ID             int     `gorm:"primaryKey;autoIncrement"`
	VehicleID      int     `gorm:"not null"`
	PricePerMinute float64 `gorm:"type:numeric(10,2);not null"`
	PricePerHour   float64 `gorm:"type:numeric(10,2);not null"`
	PricePerDay    float64 `gorm:"type:numeric(10,2);not null"`
	InsurancePrice float64 `gorm:"type:numeric(10,2)"`
	CleaningFee    float64 `gorm:"type:numeric(10,2)"`
	StartDate      time.Time
	EndDate        time.Time
}
