package router

import (
	"github.com/edaywalid/planify-backend/internal/di"
	"github.com/gin-gonic/gin"
)

func SetupRouter(container *di.Container) *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "https://devfest-batna24-frontend.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	NewPingRouter(container).SetupRouter(router)
	NewAuthRouter(container).SetupRouter(router)
	NewSwaggerRouter(container).SetupRouter(router)
	return router
}
