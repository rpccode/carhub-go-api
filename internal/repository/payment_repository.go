package repository

import (
	"carhub/internal/model"
	"errors"

	"gorm.io/gorm"
)

// Interface que define los métodos del repositorio
type PaymentRepository interface {
	Create(payment *model.PaymentToFacility) error
	FindByID(id int) (*model.PaymentToFacility, error)
	Update(payment *model.PaymentToFacility) error
	Delete(id int) error
}

type paymentRepository struct {
	db *gorm.DB
}

// NewPaymentRepository crea una nueva instancia del repositorio
func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{
		db: db,
	}
}

// Implementación de los métodos del repositorio
func (r *paymentRepository) Create(payment *model.PaymentToFacility) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) FindByID(id int) (*model.PaymentToFacility, error) {
	var payment model.PaymentToFacility
	result := r.db.First(&payment, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("payment not found")
	}
	return &payment, result.Error
}

func (r *paymentRepository) Update(payment *model.PaymentToFacility) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepository) Delete(id int) error {
	return r.db.Delete(&model.PaymentToFacility{}, id).Error
}
