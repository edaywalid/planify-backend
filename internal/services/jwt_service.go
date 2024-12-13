package services

import (
	"time"

	"github.com/edaywalid/devfest-batna24-backend/internal/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JwtService struct {
	config *config.Config
}

func NewJwtService(config *config.Config) *JwtService {
	return &JwtService{config}
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

func (s *JwtService) generateToken(userID uuid.UUID, live_time time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(live_time).Unix(),
	})

	tokenString, err := token.SignedString([]byte(s.config.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JwtService) GenerateToken(userID uuid.UUID) (*Token, error) {
	accessToken, err := s.generateToken(userID, time.Minute*15)
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: accessToken,
	}, nil
}

func (s *JwtService) ValidateToken(tokenString string) (uuid.UUID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.config.JWT_SECRET), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, jwt.ValidationError{
			Errors: jwt.ValidationErrorExpired,
		}
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, jwt.ValidationError{
			Errors: jwt.ValidationErrorClaimsInvalid,
		}
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return uuid.Nil, jwt.ValidationError{
			Errors: jwt.ValidationErrorClaimsInvalid,
		}
	}

	return uuid.Parse(userID)
}
