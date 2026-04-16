package routes

import (
	"learning-backend/container"
	"learning-backend/middleware"
	//"learning-backend/middleware"

	"github.com/gofiber/fiber/v2"
)




func  SetupUserRoutes(app *fiber.App, c *container.Dependency) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend is running")
	})

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy")
	})

	// public routes
  app.Post("/signup", c.UserHandler.SignUp) 
	app.Post("/signin", c.UserHandler.SignIn)
	app.Get("/:shopSlug", c.UserHandler.GetShopBySlug)



	// private routes
  priRoutes := app.Group("/user", middleware.AuthMiddleware(c.Config.JWTSecret)) // Apply auth middleware to all routes in this group

	priRoutes.Get("/verify", c.UserHandler.GetVerificationCode)	
	priRoutes.Post("/verify", c.UserHandler.Verify)	

	priRoutes.Post("/profile", c.UserHandler.CreateProfile)	
	priRoutes.Get("/profile", c.UserHandler.GetProfile)	
	priRoutes.Patch("/profile", c.UserHandler.UpdateProfile)	

	priRoutes.Post("/cart", c.UserHandler.AddtoCart)	
	priRoutes.Get("/cart", c.UserHandler.GetCart)	

	priRoutes.Get("/order", c.UserHandler.GetOrders)	
	priRoutes.Get("/order/:id", c.UserHandler.GetOrder)	

	priRoutes.Post("/Become-seller", c.UserHandler.BecomeSeller)	






		
	


}

