package router

import (
	"github.com/edaywalid/planify-backend/internal/di"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type SwaggerRouter struct {
	container *di.Container
}

func NewSwaggerRouter(container *di.Container) *SwaggerRouter {
	return &SwaggerRouter{container: container}
}

func (sr *SwaggerRouter) SetupRouter(r *gin.Engine) {
	r.StaticFile("/docs/swagger.yaml", "./docs/swagger.yaml")
	swaggerGroup := r.Group("/swagger")
	{
		swaggerGroup.GET("/*any", ginSwagger.CustomWrapHandler(
			&ginSwagger.Config{
				URL: "/docs/swagger.yaml",
			},
			swaggerFiles.Handler,
		))
	}

}
