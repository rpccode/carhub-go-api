package repository

import (
	"carhub/internal/model"

	"gorm.io/gorm"
)

// ModelRepository maneja las operaciones de base de datos para Model
type ModelRepository struct {
	DB *gorm.DB
}

// NewModelRepository crea un nuevo ModelRepository
func NewModelRepository(db *gorm.DB) *ModelRepository {
	return &ModelRepository{DB: db}
}

// Create guarda un nuevo modelo en la base de datos
func (r *ModelRepository) Create(model *model.Model) error {
	return r.DB.Create(model).Error
}

// GetAll obtiene todos los modelos de la base de datos
func (r *ModelRepository) GetAll() ([]model.Model, error) {
	var models []model.Model
	if err := r.DB.Find(&models).Error; err != nil {
		return nil, err
	}
	return models, nil
}
