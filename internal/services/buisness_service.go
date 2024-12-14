package services

import (
	"github.com/edaywalid/devfest-batna24-backend/internal/models"
	"github.com/edaywalid/devfest-batna24-backend/internal/repositories"
	"github.com/google/uuid"
)

type BusinessService struct {
	businessRepo *repositories.BusinessRepository
}

func NewBusinessService(businessRepo *repositories.BusinessRepository) *BusinessService {
	return &BusinessService{
		businessRepo,
	}
}

func (bs *BusinessService) CreateBusiness(business *models.Business) (*models.Business, error) {
	return bs.businessRepo.Create(business)
}

func (bs *BusinessService) GetBusinessById(id uuid.UUID) (*models.Business, error) {
	return bs.businessRepo.FindById(id)
}

func (bs *BusinessService) GetAllBusinesses(user_id uuid.UUID) ([]*models.Business, error) {
	return bs.businessRepo.FindAll(user_id)
}

func (bs *BusinessService) DeleteBusiness(id uuid.UUID) error {
	return bs.businessRepo.Delete(id)
}
