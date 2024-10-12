package model

type GeneralSettings struct {
	ID              int    `gorm:"primaryKey;autoIncrement"`
	DefaultCurrency string `gorm:"type:varchar(50);default:'Dominican Peso (DOP)'"`
	Language        string `gorm:"type:varchar(50);default:'Spanish'"`
}
