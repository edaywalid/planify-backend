package router

import (
	"github.com/edaywalid/planify-backend/internal/di"
	"github.com/gin-gonic/gin"
)

type PingRouter struct {
	container *di.Container
}

func NewPingRouter(container *di.Container) *PingRouter {
	return &PingRouter{container: container}
}

func (pr *PingRouter) SetupRouter(r *gin.Engine) {
	pingGroup := r.Group("/ping")
	{
		pingGroup.GET("", pr.Ping)
	}
}

func (pr *PingRouter) Ping(c *gin.Context) {
	pr.container.Handlers.PingHandler.Ping(c)
}
