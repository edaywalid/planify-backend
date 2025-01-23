package handlers

import (
	"net/http"

	"github.com/edaywalid/planify-backend/internal/services"
	logger "github.com/edaywalid/planify-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	logger      *logger.MyLogger
	authService *services.AuthService
}

func NewAuthHandler(
	logger *logger.MyLogger,
	authService *services.AuthService,
) *AuthHandler {
	return &AuthHandler{
		logger,
		authService,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		FullName string `json:"fullName" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.LogError().Msgf("failed binding register body : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.authService.Register(input.FullName, input.Email, input.Password)
	if err != nil {
		h.logger.LogError().Msgf("failed registering user : %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.SetCookie("accessToken", token.AccessToken, 3600, "/", "devfest-batna24-backend.onrender.com", true, true)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.LogError().Msgf("failed binding register body : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := h.authService.Login(input.Email, input.Password)
	if err != nil {
		h.logger.LogError().Msgf("failed login user : %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("acessToken", token.AccessToken, 3600, "/", "devfest-batna24-backend.onrender.com", true, true)
	c.JSON(http.StatusCreated, gin.H{"message": "User logged in successfully"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("accessToken", "", -1, "/", "devfest-batna24-backend.onrender.com", true, true)
	c.JSON(http.StatusCreated, gin.H{"message": "User logged out successfully"})
}
