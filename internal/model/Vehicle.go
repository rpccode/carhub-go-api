package model

import (
	"time"

	"github.com/jackc/pgtype"
)

type Vehicle struct {
	ID           int          `gorm:"primaryKey;autoIncrement"`
	LicensePlate string       `gorm:"type:varchar(10);not null"`
	BrandID      int          `gorm:"not null"`
	ModelID      int          `gorm:"not null"`
	Location     pgtype.Point `gorm:"type:point"`
	FuelTypeID   int          `gorm:"not null"`
	Status       string       `gorm:"type:varchar(100);not null"`
	Seats        int          `gorm:"default:4"`
	Rating       float64      `gorm:"type:numeric(2,1)"`
	IsEconomy    bool         `gorm:"default:false"`
	IsLuxury     bool         `gorm:"default:false"`
	CreatedAt    time.Time    `gorm:"autoCreateTime"`
	UpdatedAt    time.Time    `gorm:"autoUpdateTime"`
}
