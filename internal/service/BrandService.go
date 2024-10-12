package service

import (
	"carhub/internal/model"
	"carhub/internal/repository"
)

// BrandService contiene la lógica de negocio para Brand
type BrandService struct {
	Repo *repository.BrandRepository
}

// NewBrandService crea un nuevo BrandService
func NewBrandService(repo *repository.BrandRepository) *BrandService {
	return &BrandService{Repo: repo}
}

// Create crea una nueva marca
func (s *BrandService) Create(brand *model.Brand) error {
	return s.Repo.Create(brand)
}

// GetAll obtiene todas las marcas
func (s *BrandService) GetAll() ([]model.Brand, error) {
	return s.Repo.GetAll()
}
