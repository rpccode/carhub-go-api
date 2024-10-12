// controller/payment_plan_controller.go
package controller

import (
	"carhub/internal/model"
	"carhub/internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// PaymentPlanController maneja las solicitudes HTTP para PaymentPlan
type PaymentPlanController struct {
	service service.PaymentPlanService
}

// NewPaymentPlanController crea un nuevo controlador de PaymentPlan
func NewPaymentPlanController(service service.PaymentPlanService) *PaymentPlanController {
	return &PaymentPlanController{service: service}
}

// CreatePaymentPlan maneja la creación de un nuevo plan de pago
func (c *PaymentPlanController) CreatePaymentPlan(w http.ResponseWriter, r *http.Request) {
	var plan model.PaymentPlan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := c.service.CreatePaymentPlan(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(plan)
}

// GetPaymentPlanByID maneja la obtención de un plan de pago por su ID
func (c *PaymentPlanController) GetPaymentPlanByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	plan, err := c.service.GetPaymentPlanByID(id)
	if err != nil {
		http.Error(w, "Plan not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(plan)
}

// UpdatePaymentPlan maneja la actualización de un plan de pago
func (c *PaymentPlanController) UpdatePaymentPlan(w http.ResponseWriter, r *http.Request) {
	var plan model.PaymentPlan
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := c.service.UpdatePaymentPlan(&plan); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(plan)
}

// DeletePaymentPlan maneja la eliminación de un plan de pago por su ID
func (c *PaymentPlanController) DeletePaymentPlan(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := c.service.DeletePaymentPlan(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
