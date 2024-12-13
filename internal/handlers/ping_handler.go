package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingHandler struct{}

func NewPingHandler() *PingHandler {
	return &PingHandler{}
}

func (ph *PingHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "Pong")
}
