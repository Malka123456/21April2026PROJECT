package routes

import (
	"learning-backend/container"
	"learning-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupCatalogRoutes(app *fiber.App, cont *container.Container) {

// public
// listing products and categories
	app.Get("/products", cont.CatalogHandler.GetProducts)
	app.Get("/products/:id", cont.CatalogHandler.GetProduct)
	app.Get("/catagories", cont.CatalogHandler.GetCategories)
	app.Get("/catagories/:id", cont.CatalogHandler.GetCategoryById)

// private
	// manage products and categories
	selRoutes := app.Group("/seller", middleware.AuthorizeSeller) // later i will customised it

	
	// Categories
	selRoutes.Post("/categories", cont.CatalogHandler.CreateCategories)
	selRoutes.Patch("/categories/:id", cont.CatalogHandler.EditCategory)
	selRoutes.Delete("/categories/:id", cont.CatalogHandler.DeleteCategory)

	// Products
	selRoutes.Post("/products", cont.CatalogHandler.CreateProducts)
	selRoutes.Get("/products", cont.CatalogHandler.GetProducts)
	selRoutes.Get("/products/:id", cont.CatalogHandler.GetProduct)
	selRoutes.Put("/products/:id", cont.CatalogHandler.EditProduct)
	selRoutes.Patch("/products/:id", cont.CatalogHandler.UpdateStock) // update stock
	selRoutes.Delete("/products/:id", cont.CatalogHandler.DeleteProduct)


}