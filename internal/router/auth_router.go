package router

import (
	"github.com/edaywalid/devfest-batna24-backend/internal/di"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	container *di.Container
}

func NewAuthRouter(container *di.Container) *AuthRouter {
	return &AuthRouter{container: container}
}

func (ar *AuthRouter) SetupRouter(r *gin.Engine) {
	authGroup := r.Group("/api/v1")
	{
		authGroup.POST("/login", ar.Login)
		authGroup.POST("/register", ar.Register)
	}

}

func (ar *AuthRouter) Login(c *gin.Context) {
	ar.container.Handlers.AuthHandler.Login(c)
}

func (ar *AuthRouter) Register(c *gin.Context) {
	ar.container.Handlers.AuthHandler.Register(c)
}
