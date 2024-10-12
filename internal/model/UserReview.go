package model

import "time"

type UserReview struct {
	ID        int       `gorm:"primaryKey;autoIncrement"`
	VehicleID int       `gorm:"not null"`
	UserID    int       `gorm:"not null"`
	Rating    float64   `gorm:"type:numeric(2,1);not null"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
