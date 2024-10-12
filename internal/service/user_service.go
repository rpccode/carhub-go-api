package service

import (
	"carhub/internal/model"
	"carhub/internal/repository"
)

type UserService interface {
	CreateUser(user *model.User) error
	GetAllUsers() ([]model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) CreateUser(user *model.User) error {
	return s.userRepo.Create(user)
}

func (s *userService) GetAllUsers() ([]model.User, error) {
	return s.userRepo.FindAll()
}
