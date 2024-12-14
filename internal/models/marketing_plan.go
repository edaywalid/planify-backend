package models

import (
	"time"

	"github.com/google/uuid"
)

type MarketingPlan struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID     uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name       string    `gorm:"not null" json:"name"`
	BusinessID uuid.UUID `gorm:"type:uuid;not null" json:"business_id"`
	Business   Business  `gorm:"foreignKey:BusinessID;constraint:OnDelete:CASCADE;" json:"business"`
	Phases     []Phase   `gorm:"foreignKey:MarketingPlanID;constraint:OnDelete:CASCADE;" json:"phases"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Phase struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	MarketingPlanID uuid.UUID `gorm:"type:uuid;not null" json:"marketing_plan_id"`
	Title           string    `gorm:"not null" json:"title"`
	Steps           []Step    `gorm:"foreignKey:PhaseID;constraint:OnDelete:CASCADE;" json:"steps"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Step struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	PhaseID     uuid.UUID `gorm:"type:uuid;not null" json:"phase_id"`
	Title       string    `gorm:"not null" json:"title"`
	Description string    `gorm:"not null" json:"description"`
	Timeline    string    `json:"timeline,omitempty"`
	Resources   string    `json:"resources,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
