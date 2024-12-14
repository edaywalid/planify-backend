package services

import (
	"github.com/edaywalid/devfest-batna24-backend/internal/models"
	"github.com/edaywalid/devfest-batna24-backend/internal/repositories"
	"github.com/google/uuid"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{
		userRepo,
	}
}

func (us *UserService) GetUserById(id uuid.UUID) (*models.User, error) {
	return us.userRepo.FindById(id)
}

func (us *UserService) DeleteUser(id uuid.UUID) error {
	return us.userRepo.Delete(id)
}
