package services

import (
	"errors"

	"github.com/edaywalid/devfest-batna24-backend/internal/models"
	"github.com/edaywalid/devfest-batna24-backend/internal/repositories"
	"github.com/edaywalid/devfest-batna24-backend/pkg/utils"
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

func (s *AuthService) Register(fullname, email, password string) error {
	if !utils.ValidatePassword(password) {
		return errors.New("password not secure")
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		FullName: fullname,
		Email:    email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(username, password string) (*Token, error) {
	user, err := s.userRepo.FindByUsername(username)
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
