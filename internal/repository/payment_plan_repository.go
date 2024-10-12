// repository/payment_plan_repository.go
package repository

import (
	"carhub/internal/model"
	"errors"

	"gorm.io/gorm"
)

// PaymentPlanRepository define la interfaz para las operaciones CRUD
type PaymentPlanRepository interface {
	Create(plan *model.PaymentPlan) error
	FindByID(id int) (*model.PaymentPlan, error)
	Update(plan *model.PaymentPlan) error
	Delete(id int) error
}

// paymentPlanRepositoryImpl es la implementación de PaymentPlanRepository
type paymentPlanRepositoryImpl struct {
	db *gorm.DB
}

// NewPaymentPlanRepository crea un nuevo repositorio de PaymentPlan
func NewPaymentPlanRepository(db *gorm.DB) PaymentPlanRepository {
	return &paymentPlanRepositoryImpl{db: db}
}

func (r *paymentPlanRepositoryImpl) Create(plan *model.PaymentPlan) error {
	return r.db.Create(plan).Error
}

func (r *paymentPlanRepositoryImpl) FindByID(id int) (*model.PaymentPlan, error) {
	var plan model.PaymentPlan
	if err := r.db.First(&plan, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return &plan, nil
}

func (r *paymentPlanRepositoryImpl) Update(plan *model.PaymentPlan) error {
	return r.db.Save(plan).Error
}

func (r *paymentPlanRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&model.PaymentPlan{}, id).Error
}
