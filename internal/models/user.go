package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	FullName string    `json:"fullName"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}
