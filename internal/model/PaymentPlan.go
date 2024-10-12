package model

import (
	"time"
)

type PaymentPlan struct {
	ID              int       `gorm:"primaryKey;autoIncrement"`
	ReservationID   int       `gorm:"not null"`
	TotalAmount     float64   `gorm:"type:numeric(10,2)"`
	InitialPayment  float64   `gorm:"type:numeric(10,2)"`
	RemainingAmount float64   `gorm:"type:numeric(10,2)"`
	NumInstallments int       `gorm:"not null"`
	InterestRate    float64   `gorm:"type:numeric(5,2)"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
}
