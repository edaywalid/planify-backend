package router

import (
	"github.com/edaywalid/planify-backend/internal/di"
	"github.com/gin-gonic/gin"
)

func SetupRouter(container *di.Container) *gin.Engine {
	router := gin.Default()
	router.Use(container.Middlewares.CorsMiddleWare.CORSMiddleware())
	NewPingRouter(container).SetupRouter(router)
	NewAuthRouter(container).SetupRouter(router)
	NewSwaggerRouter(container).SetupRouter(router)
	return router
}
