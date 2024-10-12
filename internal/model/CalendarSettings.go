package model

import (
	"time"
)

type CalendarSettings struct {
	ID           int           `gorm:"primaryKey;autoIncrement"`
	ReminderTime time.Duration `gorm:"default:interval '1 day'"`
	DefaultView  string        `gorm:"type:varchar(20);default:'Day'"`
	ShowWeekends bool          `gorm:"default:true"`
}
