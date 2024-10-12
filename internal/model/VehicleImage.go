package model

import "time"

type VehicleImage struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	VehicleID int       `gorm:"not null"`
	ImageURL  string    `gorm:"type:varchar(255);not null"`
	IsPrimary bool      `gorm:"default:false"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
