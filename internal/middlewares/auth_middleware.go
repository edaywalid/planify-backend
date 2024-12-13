package middlewares

import (
	"github.com/edaywalid/devfest-batna24-backend/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthMiddleware struct {
	jwtService *services.JwtService
}

func NewAuthMiddleware(jwtService *services.JwtService) *AuthMiddleware {
	return &AuthMiddleware{jwtService}
}

func (m *AuthMiddleware) AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("accessToken")
		if err != nil {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		userID, err := m.jwtService.ValidateToken(accessToken)
		if err != nil {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		ctx.Set("user_id", userID)
		ctx.Next()
	}
}
