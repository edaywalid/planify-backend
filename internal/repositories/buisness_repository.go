package repositories

import (
	"github.com/edaywalid/devfest-batna24-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessRepository struct {
	db *gorm.DB
}

func NewBuisnessRepository(db *gorm.DB) *BusinessRepository {
	return &BusinessRepository{
		db,
	}
}

func (br *BusinessRepository) Create(buisness *models.Business) (*models.Business, error) {
	var inserted_buisness models.Business
	err := br.db.Create(buisness).First(&inserted_buisness).Error
	return &inserted_buisness, err
}

func (br *BusinessRepository) FindById(id uuid.UUID) (*models.Business, error) {
	var business models.Business
	err := br.db.Where("id = ?", id).First(&business).Error
	return &business, err
}

func (br *BusinessRepository) FindAll(user_id uuid.UUID) ([]*models.Business, error) {
	var businesses []*models.Business
	err := br.db.Where("user_id = ?", user_id).Find(&businesses).Error
	return businesses, err
}

func (br *BusinessRepository) Delete(id uuid.UUID) error {
	err := br.db.Where("id = ?", id).Delete(&models.Business{}).Error
	return err
}
