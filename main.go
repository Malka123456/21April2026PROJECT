package main

import (
	"learning-backend/container"
	"learning-backend/routes"

	//"learning-backend/handlers"
	//"learning-backend/models"

	"fmt"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)




func main() {


	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	cont := container.BuildContainer()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))

	setupRoutes(app, cont)

	//config.LoadKeys()


	fmt.Println("server is running")

  log.Fatal(app.Listen(":" + cont.Config.AppPort))

}

func setupRoutes(app *fiber.App, cont *container.Dependency) {

	routes.SetupUserRoutes(app, cont)
	routes.SetupCatalogRoutes(app, cont)
	
}