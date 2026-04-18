package routes

import (
	"learning-backend/container"
	"learning-backend/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupCatalogRoutes(app *fiber.App, cont *container.Dependency) {

// public
// listing products and categories
	app.Get("/products", cont.CatalogHandler.GetProductsForPublic)
	app.Get("/products/:id", cont.CatalogHandler.GetProductForPublic)
	app.Get("/categories", cont.CatalogHandler.GetCategories)
	app.Get("/categories/:id", cont.CatalogHandler.GetCategoryById)

	// product url   
  app.Get("/shops/:shopSlug/:productSlug", cont.CatalogHandler.GetProductBySlug) 


  // private
	// manage products and categories
	selRoutes := app.Group("/seller", middleware.AuthMiddleware(cont.Config.JWTSecret)) // later i will customised it

	
	// Categories
	selRoutes.Post("/categories", cont.CatalogHandler.CreateCategories)
	selRoutes.Patch("/categories/:id", cont.CatalogHandler.EditCategory)
	selRoutes.Delete("/categories/:id", cont.CatalogHandler.DeleteCategory)

	// Products
	selRoutes.Post("/products", cont.CatalogHandler.CreateProducts)
	selRoutes.Get("/products", cont.CatalogHandler.GetProductsForSeller)
	selRoutes.Get("/products/:id", cont.CatalogHandler.GetProductForSeller)
	selRoutes.Put("/products/:id", cont.CatalogHandler.EditProduct)
	selRoutes.Patch("/products/:id", cont.CatalogHandler.UpdateStock) // update stock
	selRoutes.Delete("/products/:id", cont.CatalogHandler.DeleteProduct)


	
}