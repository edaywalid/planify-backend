package di

import (
	"context"
	"time"

	"github.com/edaywalid/devfest-batna24-backend/internal/config"
	"github.com/edaywalid/devfest-batna24-backend/internal/db"
	"github.com/edaywalid/devfest-batna24-backend/internal/handlers"
	"github.com/edaywalid/devfest-batna24-backend/internal/middlewares"
	"github.com/edaywalid/devfest-batna24-backend/internal/repositories"
	"github.com/edaywalid/devfest-batna24-backend/internal/services"
	logger "github.com/edaywalid/devfest-batna24-backend/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Container struct {
	Services     *Services
	Repositories *Repositories
	Handlers     *Handlers
	Config       *config.Config
	Databases    *Databases
	Middlewares  *Middlewares
	Logger       *logger.MyLogger
	Cache        *Cache
}

type (
	Services struct {
		AuthService *services.AuthService
		JwtService  *services.JwtService
	}

	Repositories struct {
		UserRepository *repositories.UserRepository
	}
	Handlers struct {
		PingHandler    *handlers.PingHandler
		SwaggerHandler *handlers.SwaggerHandler
		AuthHandler    *handlers.AuthHandler
	}
	Databases struct {
		postgres *gorm.DB
		mongo    *mongo.Client
	}
	Cache struct {
		// redis *cache.Redis
	}
	Middlewares struct {
		AuthMiddleWare *middlewares.AuthMiddleware
		CorsMiddleWare *middlewares.CorsMiddleware
	}
)

func NewContainer(log *logger.MyLogger) (*Container, error) {
	var container Container
	container.Logger = log

	container.Logger.LogInfo().Msg("Loading config")
	config, err := config.LoadConfig()
	if err != nil {
		container.Logger.LogError().Msgf("Error loading config: %v", err)
		return nil, err
	}
	container.Logger.LogInfo().Msg("Config loaded successfully")
	container.Config = config

	container.Logger.LogInfo().Msg("Initializing databases")
	if err := container.initDatabases(); err != nil {
		log.LogError().Msgf("Error initializing databases: %v", err)
		return nil, err
	}
	log.LogInfo().Msg("Databases initialized successfully")

	log.LogInfo().Msg("Initializing services, repositories, and handlers")

	container.InitRepositories()
	log.LogInfo().Msg("Repositories initialized successfully")

	container.InitServices()
	log.LogInfo().Msg("Services initialized successfully")

	container.InitHandlers()
	log.LogInfo().Msg("Handlers initialized successfully")

	container.InitMiddlewares()
	log.LogInfo().Msg("Middlewared initialized scuccessfully")
	return &container, nil
}

func (c *Container) initDatabases() error {
	postgres, err := db.InitPSQL(c.Config)
	if err != nil {
		return err
	}

	mongo, err := db.InitMongo(c.Config)
	if err != nil {
		return err
	}

	c.Databases = &Databases{
		postgres: postgres,
		mongo:    mongo,
	}
	return nil
}

func (c *Container) InitServices() {
	c.Services = &Services{
		JwtService: services.NewJwtService(c.Config),
		// RedisService: services.NewRedisService(c.Config),
	}
	c.Services.AuthService = services.NewAuthService(
		c.Repositories.UserRepository,
		c.Services.JwtService,
	)
}

func (c *Container) InitCache() {
	// redis, err := cache.NewRedis(c.Config.REDIS_ADDR)
	// if err != nil {
	// 	c.Logger.LogError().Msgf("Error initializing redis: %v", err)
	// 	return
	// }
	// c.Cache = &Cache{
	// 	redis: redis,
	// }
}

func (c *Container) InitRepositories() {
	c.Repositories = &Repositories{
		UserRepository: repositories.NewUserRepository(c.Databases.postgres),
	}
}

func (c *Container) InitHandlers() {
	handlers := &Handlers{
		PingHandler: handlers.NewPingHandler(),
		SwaggerHandler: handlers.NewSwaggerHandler(
			c.Config,
			c.Logger,
		),
		AuthHandler: handlers.NewAuthHandler(
			c.Logger,
			c.Services.AuthService,
		),
	}
	c.Handlers = handlers
}

func (c *Container) InitMiddlewares() {
	c.Middlewares = &Middlewares{
		AuthMiddleWare: middlewares.NewAuthMiddleware(c.Services.JwtService),
		CorsMiddleWare: middlewares.NewCorsMiddleware(),
	}
}

func (c *Container) Close() error {
	c.Logger.LogInfo().Msg("Closing databases")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if c.Databases.mongo != nil {
		c.Databases.mongo.Disconnect(ctx)
	}

	return nil
}
