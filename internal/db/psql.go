package db

import (
	"github.com/edaywalid/planify-backend/internal/config"
	"github.com/edaywalid/planify-backend/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPSQL(config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.POSTGRES_URI), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.User{}, &models.Business{}, &models.MarketingPlan{}, &models.Phase{}, &models.Step{}, &models.User{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
