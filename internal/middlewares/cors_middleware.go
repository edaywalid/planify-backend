package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type CorsMiddleware struct{}

func NewCorsMiddleware() *CorsMiddleware {
	return &CorsMiddleware{}
}

func isAllowedOrigin(origin string, allowedOrigins []string) bool {
	for _, allowed := range allowedOrigins {
		if origin == allowed {
			return true
		}
	}
	return false
}

func (m *CorsMiddleware) CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		env := os.Getenv("ENV")

		var allowedOrigins []string
		if env == "production" {
			allowedOrigins = strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ",")
		} else {
			allowedOrigins = []string{"*"}
		}

		origin := c.GetHeader("Origin")

		if env != "production" || isAllowedOrigin(origin, allowedOrigins) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			return
		}

		c.Next()
	}
}
