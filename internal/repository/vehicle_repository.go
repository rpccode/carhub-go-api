// repository/vehicle_repository.go
package repository

import (
	"carhub/internal/model"

	"gorm.io/gorm"
)

type VehicleRepository interface {
	Create(vehicle *model.Vehicle) error
	FindByID(id int) (*model.Vehicle, error)
	Update(vehicle *model.Vehicle) error
	Delete(id int) error
}

type vehicleRepositoryImpl struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepositoryImpl{db: db}
}

func (r *vehicleRepositoryImpl) Create(vehicle *model.Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *vehicleRepositoryImpl) FindByID(id int) (*model.Vehicle, error) {
	var vehicle model.Vehicle
	if err := r.db.First(&vehicle, id).Error; err != nil {
		return nil, err
	}
	return &vehicle, nil
}

func (r *vehicleRepositoryImpl) Update(vehicle *model.Vehicle) error {
	return r.db.Save(vehicle).Error
}

func (r *vehicleRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&model.Vehicle{}, id).Error
}
