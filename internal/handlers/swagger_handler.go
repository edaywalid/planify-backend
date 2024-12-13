package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/edaywalid/devfest-batna24-backend/internal/config"
	logger "github.com/edaywalid/devfest-batna24-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	httpSwagger "github.com/swaggo/http-swagger"
)

type SwaggerHandler struct {
	config *config.Config
	logger *logger.MyLogger
}

func NewSwaggerHandler(
	config *config.Config,
	logger *logger.MyLogger,
) *SwaggerHandler {
	return &SwaggerHandler{
		config: config,
		logger: logger,
	}
}
func (sh *SwaggerHandler) ServeYamlDocs(c *gin.Context) {
	cwd, err := os.Getwd()
	if err != nil {
		sh.logger.LogError().Err(err).Msg("Failed to get current working directory")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	yamlFile := filepath.Join(cwd, "docs", "swagger.yaml")
	if _, err := os.Stat(yamlFile); os.IsNotExist(err) {
		sh.logger.LogError().Err(err).Msg("Swagger file not found")
		c.JSON(http.StatusNotFound, gin.H{"error": "Swagger file not found"})
		return
	}

	sh.logger.LogInfo().Msg("Serving swagger file")
	c.Header("Content-Type", "application/x-yaml")
	c.File(yamlFile)
}

func (sh *SwaggerHandler) ServeSwaggerUI() http.Handler {
	url := "http://localhost:" + sh.config.PORT + "/swagger/doc.yaml"
	if sh.config.IsProduction() {
		url = sh.config.DOCS_URL
	}

	return httpSwagger.Handler(
		httpSwagger.URL(url),
	)
}
