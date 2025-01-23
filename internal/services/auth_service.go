package services

import (
	"errors"

	"github.com/edaywalid/planify-backend/internal/models"
	"github.com/edaywalid/planify-backend/internal/repositories"
	"github.com/edaywalid/planify-backend/pkg/utils"
)

type Token struct {
	AccessToken string
}

type AuthService struct {
	userRepo   *repositories.UserRepository
	jwtService *JwtService
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	jwtService *JwtService,
) *AuthService {
	return &AuthService{
		userRepo,
		jwtService,
	}
}

func (s *AuthService) Register(fullname, email, password string) (*Token, error) {
	if !utils.ValidatePassword(password) {
		return nil, errors.New("password not secure")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		FullName: fullname,
		Email:    email,
		Password: string(hashedPassword),
	}

	inserted_user, err := s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	token, err := s.jwtService.GenerateToken(inserted_user.ID)
	if err != nil {
		return nil, err
	}

	return token, err
}

func (s *AuthService) Login(email, password string) (*Token, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPasswordHash(password, user.Password)
	if err != nil {
		return nil, err
	}

	token, err := s.jwtService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}
	return token, nil
}
