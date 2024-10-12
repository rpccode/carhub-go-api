package model

type User struct {
	ID        int    `gorm:"primaryKey;autoIncrement"`
	Username  string `gorm:"type:varchar(255);not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Email     string `gorm:"type:varchar(255);not null"`
	Phone     string `gorm:"type:varchar(15)"`
	PushToken string `gorm:"type:text"`
}
