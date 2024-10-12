package service

import (
	"carhub/internal/model"
	"carhub/internal/repository"
	"errors"
	"time"
)

// Interfaz para el servicio
type PaymentService interface {
	CreatePayment(payment *model.PaymentToFacility) error
	GetPaymentByID(id int) (*model.PaymentToFacility, error)
	UpdatePayment(payment *model.PaymentToFacility) error
	DeletePayment(id int) error
	MarkPaymentAsPaid(id int, paidAt time.Time) error
}

type paymentService struct {
	repo repository.PaymentRepository
}

// NewPaymentService crea una nueva instancia del servicio
func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentService{
		repo: repo,
	}
}

// Implementación de los métodos del servicio
func (s *paymentService) CreatePayment(payment *model.PaymentToFacility) error {
	if payment.DueDate.Before(time.Now()) {
		return errors.New("due date cannot be in the past")
	}
	return s.repo.Create(payment)
}

func (s *paymentService) GetPaymentByID(id int) (*model.PaymentToFacility, error) {
	return s.repo.FindByID(id)
}

func (s *paymentService) UpdatePayment(payment *model.PaymentToFacility) error {
	return s.repo.Update(payment)
}

func (s *paymentService) DeletePayment(id int) error {
	return s.repo.Delete(id)
}

func (s *paymentService) MarkPaymentAsPaid(id int, paidAt time.Time) error {
	payment, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}

	if payment.Status == "paid" {
		return errors.New("payment already marked as paid")
	}

	payment.PaidAt = paidAt
	payment.Status = "paid"

	return s.repo.Update(payment)
}
