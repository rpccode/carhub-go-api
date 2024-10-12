package model

import "time"

type PaymentToFacility struct {
	ID            int       `gorm:"primaryKey;autoIncrement"`
	PaymentPlanID int       `gorm:"not null"`
	DueDate       time.Time `gorm:"type:date"`
	OverdueAmount float64   `gorm:"type:numeric(10,2)"`
	PaidAmount    float64   `gorm:"type:numeric(10,2)"`
	PaidAt        time.Time `gorm:"type:timestamp"`
	Status        string    `gorm:"type:varchar(20);not null"`
}
