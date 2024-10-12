package service

import (
	"carhub/internal/model"
	"carhub/internal/repository"
)

// Importa tu repositorio

// ModelService contiene la lógica de negocio para Model
type ModelService struct {
	Repo *repository.ModelRepository
}

// NewModelService crea un nuevo ModelService
func NewModelService(repo *repository.ModelRepository) *ModelService {
	return &ModelService{Repo: repo}
}

// Create crea un nuevo modelo
func (s *ModelService) Create(model *model.Model) error {
	return s.Repo.Create(model)
}

// GetAll obtiene todos los modelos
func (s *ModelService) GetAll() ([]model.Model, error) {
	return s.Repo.GetAll()
}
