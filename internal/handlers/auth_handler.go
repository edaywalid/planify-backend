package handlers

import (
	"net/http"
	"time"

	"github.com/edaywalid/devfest-batna24-backend/internal/services"
	logger "github.com/edaywalid/devfest-batna24-backend/pkg/utils"
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

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("accessToken", token.AccessToken, 60*60*24*30, "/", "localhost", false, true)
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
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    token.AccessToken,
		Path:     "/",
		Domain:   "devfest-batna24-backend.onrender.com",
		Expires:  time.Now().Add(30 * 24 * time.Hour),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})
	c.JSON(http.StatusCreated, gin.H{"message": "User logged in successfully"})
}

func (h *AuthHandler) Logout(c *gin.Context) {
	domain := "devfest-batna24-backend.onrender.com"
	http.SetCookie(c.Writer, &http.Cookie{
		Name:     "accessToken",
		Value:    "",
		Path:     "/",
		Domain:   domain,
		Expires:  time.Unix(0, 0),
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
	})

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "User logged out successfully",
	})
}
