package middleware

import (
	//"go/token"
	//"learning-backend/config"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	//"github.com/golang-jwt/jwt/v4"
)


func AuthMiddleware(secret string) fiber.Handler {

	return func(c *fiber.Ctx) error {

	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Missing token",
		})	
	}

	parts := strings.Split(authHeader, " ")

	if len(parts) != 2 {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid format",
		})
	}

	tokenString := parts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error)  {
		return []byte(secret), nil
		})

		if err != nil || !token.Valid { 
			return c.Status(401).JSON(fiber.Map{
				"error": "Invalid token",
			})
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Locals("user_id", claims["user_id"])

		return c.Next()
	}

}

// Later real code 

func AuthorizeSeller(ctx *fiber.Ctx) error {

		return ctx.Next()
}


