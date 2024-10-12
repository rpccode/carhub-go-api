package model

type Model struct {
	ID      int    `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"type:varchar(50);not null"`
	BrandID int    `gorm:"not null"`
}
