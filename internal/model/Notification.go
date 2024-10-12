package model

import "time"

type Notification struct {
	ID      int       `gorm:"primaryKey;autoIncrement"`
	UserID  int       `gorm:"not null"`
	Message string    `gorm:"type:text;not null"`
	SentAt  time.Time `gorm:"autoCreateTime"`
	Type    string    `gorm:"type:text;not null"`
}
