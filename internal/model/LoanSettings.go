package model

type LoanSettings struct {
	ID             int     `gorm:"primaryKey;autoIncrement"`
	InterestRate   float64 `gorm:"type:numeric(5,2)"`
	LatePaymentFee float64 `gorm:"type:numeric(5,2)"`
}
