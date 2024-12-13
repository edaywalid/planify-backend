package router

import (
	"github.com/edaywalid/devfest-batna24-backend/internal/di"
	"github.com/gin-gonic/gin"
)

type SwaggerRouter struct {
	container *di.Container
}

func NewSwaggerRouter(container *di.Container) *SwaggerRouter {
	return &SwaggerRouter{container: container}
}

func (sr *SwaggerRouter) SetupRouter(r *gin.Engine) {
	authGroup := r.Group("/")
	{
		authGroup.GET("/api-docs", func(c *gin.Context) {
			sr.container.Handlers.SwaggerHandler.ServeSwaggerUI()
		})
	}
}
