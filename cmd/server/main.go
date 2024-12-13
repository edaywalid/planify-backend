package main

import (
	"github.com/edaywalid/devfest-batna24-backend/internal/di"
	"github.com/edaywalid/devfest-batna24-backend/internal/router"
	logger "github.com/edaywalid/devfest-batna24-backend/pkg/utils"
)

func main() {
	log := logger.NewLogger()

	container, err := di.NewContainer(&log)
	if err != nil {
		log.LogError().Msgf("Error creating container: %v", err)
		return
	}

	defer func() {
		log.LogInfo().Msg("Closing container")
		container.Close()
	}()
	log.LogInfo().Msg("Container created successfully")

	log.LogInfo().Msg("Setting up router")

	r := router.SetupRouter(container)
	log.LogInfo().Msg("Router setup successfully")

	log.LogInfo().Msg("Starting server")
	log.LogInfo().Msgf("Server started on port %s", container.Config.PORT)

	if err := r.Run(":" + container.Config.PORT); err != nil {
		log.LogError().Msgf("Error starting server: %v", err)
	}
}
