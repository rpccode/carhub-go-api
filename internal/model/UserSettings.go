package model

type UserSettings struct {
	UserID                 int `gorm:"primaryKey"`
	GeneralSettingsID      int
	NotificationSettingsID int
	CalendarSettingsID     int
	SecuritySettingsID     int
}
