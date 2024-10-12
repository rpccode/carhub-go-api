package model

type FuelType struct {
	ID   int    `gorm:"primaryKey;autoIncrement"`
	Type string `gorm:"type:varchar(20);not null"`
}
