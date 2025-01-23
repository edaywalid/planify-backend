package router

import (
	"github.com/edaywalid/planify-backend/internal/di"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	container *di.Container
}

func NewUserRouter(container *di.Container) *UserRouter {
	return &UserRouter{
		container: container,
	}
}

func (ur *UserRouter) SetupRouter(r *gin.Engine) {
	UserGroup := r.Group("/user", ur.container.Middlewares.AuthMiddleWare.AuthMiddleWare())
	{
		UserGroup.GET("/:id", ur.GetUserById)
		UserGroup.DELETE("/:id", ur.DeleteUser)
	}
}

func (ur *UserRouter) GetUserById(c *gin.Context) {
	ur.container.Handlers.UserHandler.GetUserById(c)
}

func (ur *UserRouter) DeleteUser(c *gin.Context) {
	ur.container.Handlers.UserHandler.GetUserById(c)
}
