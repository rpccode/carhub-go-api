// service/vehicle_service.go
package service

import (
	"carhub/internal/model"
	"carhub/internal/repository"
)

type VehicleService interface {
	CreateVehicle(vehicle *model.Vehicle) error
	GetVehicleByID(id int) (*model.Vehicle, error)
	UpdateVehicle(vehicle *model.Vehicle) error
	DeleteVehicle(id int) error
}

type vehicleServiceImpl struct {
	repo repository.VehicleRepository
}

func NewVehicleService(repo repository.VehicleRepository) VehicleService {
	return &vehicleServiceImpl{repo: repo}
}

func (s *vehicleServiceImpl) CreateVehicle(vehicle *model.Vehicle) error {
	// Validaciones o cálculos específicos pueden ir aquí
	if vehicle.Seats <= 0 {
		vehicle.Seats = 4 // Valor por defecto
	}
	return s.repo.Create(vehicle)
}

func (s *vehicleServiceImpl) GetVehicleByID(id int) (*model.Vehicle, error) {
	return s.repo.FindByID(id)
}

func (s *vehicleServiceImpl) UpdateVehicle(vehicle *model.Vehicle) error {
	return s.repo.Update(vehicle)
}

func (s *vehicleServiceImpl) DeleteVehicle(id int) error {
	return s.repo.Delete(id)
}
