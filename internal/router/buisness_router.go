package router

import (
	"github.com/edaywalid/devfest-batna24-backend/internal/di"
	"github.com/gin-gonic/gin"
)

type BuisnessRouter struct {
	container *di.Container
}

func NewBuisnessRouter(container *di.Container) *BuisnessRouter {
	return &BuisnessRouter{
		container: container,
	}
}

func (br *BuisnessRouter) SetupRouter(r *gin.Engine) {
	BuisnessGroup := r.Group("/buisness", br.container.Middlewares.AuthMiddleWare.AuthMiddleWare())
	{
		BuisnessGroup.GET("/", br.GetAllBuisnesses)
		BuisnessGroup.GET("/:id", br.GetBuisnessById)
		BuisnessGroup.POST("/", br.AddBuisness)
		BuisnessGroup.DELETE("/:id", br.DeleteBuisness)
	}
}

func (br *BuisnessRouter) GetBuisnessById(c *gin.Context) {
	br.container.Handlers.BusinessHandler.GetBusinessById(c)
}

func (br *BuisnessRouter) GetAllBuisnesses(c *gin.Context) {
	br.container.Handlers.BusinessHandler.GetAllBusinesses(c)
}

func (br *BuisnessRouter) DeleteBuisness(c *gin.Context) {
	br.container.Handlers.BusinessHandler.DeleteBusiness(c)
}

func (br *BuisnessRouter) AddBuisness(c *gin.Context) {
	br.container.Handlers.BusinessHandler.AddBusiness(c)
}
