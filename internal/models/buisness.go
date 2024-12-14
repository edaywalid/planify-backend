package models

import (
	"time"

	"github.com/google/uuid"
)

type Business struct {
	ID             uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID         uuid.UUID `gorm:"type:uuid;not null" json:"user_id"`
	Name           string    `gorm:"not null;unique" json:"name"`
	Industry       string    `gorm:"not null" json:"industry"`
	TargetAudience string    `gorm:"not null" json:"target_audience"`
	Goals          string    `gorm:"type:text" json:"goals"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
