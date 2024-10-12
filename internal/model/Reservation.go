package model

import (
	"time"
)

type Reservation struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	UserID    int       `gorm:"not null"`
	VehicleID int       `gorm:"not null"`
	StartTime time.Time `gorm:"not null"`
	EndTime   time.Time `gorm:"not null"`
	Status    string    `gorm:"type:varchar(100);not null"`
	Price     float64   `gorm:"type:numeric(10,2);not null"`
}
