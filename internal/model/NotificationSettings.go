package model

type NotificationSettings struct {
	ID                int  `gorm:"primaryKey;autoIncrement"`
	PushNotification  bool `gorm:"default:true"`
	EmailNotification bool `gorm:"default:true"`
	SMSNotification   bool `gorm:"default:true"`
}
