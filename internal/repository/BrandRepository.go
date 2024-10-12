package repository

import (
	"carhub/internal/model"

	"gorm.io/gorm"
	// Importa tu modelo
)

// BrandRepository maneja las operaciones de base de datos para Brand
type BrandRepository struct {
	DB *gorm.DB
}

// NewBrandRepository crea un nuevo BrandRepository
func NewBrandRepository(db *gorm.DB) *BrandRepository {
	return &BrandRepository{DB: db}
}

// Create guarda una nueva marca en la base de datos
func (r *BrandRepository) Create(brand *model.Brand) error {
	return r.DB.Create(brand).Error
}

// GetAll obtiene todas las marcas de la base de datos
func (r *BrandRepository) GetAll() ([]model.Brand, error) {
	var brands []model.Brand
	if err := r.DB.Find(&brands).Error; err != nil {
		return nil, err
	}
	return brands, nil
}
