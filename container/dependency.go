package container

import (
	"learning-backend/config"
	"learning-backend/database"
	"learning-backend/handlers"
	"learning-backend/helper"
	"learning-backend/repository"
	"learning-backend/service"

	"gorm.io/gorm"
)

type Container struct {
	// 🔹 Core dependencies
	DB     *gorm.DB
	Config config.AppConfig
	Auth   helper.AuthHelper

	// 🔹 Services
	UserService *service.UserService
	CatalogService *service.CatalogService

	// 🔹 Repositories
	UserRepo *repository.UserRepository
	CatalogRepo *repository.CatalogRepository

	// 🔹 Handlers
	UserHandler *handlers.UserHandler
	CatalogHandler *handlers.CatalogHandler
}

func BuildContainer() *Container {

	// 🔹 base
	config := config.LoadConfig()

	db := database.InitDB(config.DBUrl)
	database.Migrate(db)
	auth := helper.NewAuthHelper(config)

	// 🔹 repositories
	userRepo := repository.NewUserRepository(db)
	catalogRepo := repository.NewCatalogRepository(db)

	
	// 🔹 services
	userService := service.NewUserService(userRepo, auth)
		catalogService := service.NewCatalogService(auth, catalogRepo)


	// 🔹 handlers
	userHandler := handlers.NewUserHandler(userService)
		catalogHandler := handlers.NewCatalogHandler(catalogService)


	return &Container{
		DB: db,
		Config: config,
		Auth: auth,

		UserService: userService,
		UserHandler: userHandler,
		CatalogService: catalogService,
		CatalogHandler: catalogHandler,
	}
}