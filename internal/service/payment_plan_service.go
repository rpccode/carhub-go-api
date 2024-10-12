// service/payment_plan_service.go
package service

import (
	"carhub/internal/model"
	"carhub/internal/repository"
	"errors"
)

// PaymentPlanService define la interfaz para la lógica de negocio
type PaymentPlanService interface {
	CreatePaymentPlan(plan *model.PaymentPlan) error
	GetPaymentPlanByID(id int) (*model.PaymentPlan, error)
	UpdatePaymentPlan(plan *model.PaymentPlan) error
	DeletePaymentPlan(id int) error
}

// paymentPlanServiceImpl es la implementación de PaymentPlanService
type paymentPlanServiceImpl struct {
	repo repository.PaymentPlanRepository
}

// NewPaymentPlanService crea un nuevo servicio de PaymentPlan
func NewPaymentPlanService(repo repository.PaymentPlanRepository) PaymentPlanService {
	return &paymentPlanServiceImpl{repo: repo}
}

func (s *paymentPlanServiceImpl) CreatePaymentPlan(plan *model.PaymentPlan) error {
	// Lógica de negocio, como calcular intereses o validar entradas
	if plan.TotalAmount < plan.InitialPayment {
		return errors.New("Initial payment cannot be greater than total amount")
	}
	return s.repo.Create(plan)
}

func (s *paymentPlanServiceImpl) GetPaymentPlanByID(id int) (*model.PaymentPlan, error) {
	return s.repo.FindByID(id)
}

func (s *paymentPlanServiceImpl) UpdatePaymentPlan(plan *model.PaymentPlan) error {
	return s.repo.Update(plan)
}

func (s *paymentPlanServiceImpl) DeletePaymentPlan(id int) error {
	return s.repo.Delete(id)
}
