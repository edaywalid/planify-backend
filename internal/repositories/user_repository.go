package repositories

import (
	"github.com/edaywalid/planify-backend/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	var inserted_user models.User
	err := r.db.Create(user).First(&inserted_user).Error
	return &inserted_user, err
}

func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (ur *UserRepository) FindById(id uuid.UUID) (*models.User, error) {
	var user models.User
	err := ur.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (ur *UserRepository) Delete(id uuid.UUID) error {
	err := ur.db.Where("id = ?", id).Delete(&models.User{}).Error
	return err
}

func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}
