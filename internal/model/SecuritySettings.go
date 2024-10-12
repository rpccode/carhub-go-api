package model

type SecuritySettings struct {
	ID                      int  `gorm:"primaryKey;autoIncrement"`
	ChangePassword          bool `gorm:"default:true"`
	TwoFactorAuthentication bool `gorm:"default:false"`
}
