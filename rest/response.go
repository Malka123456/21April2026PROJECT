package rest

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func ErrorMessage(ctx *fiber.Ctx, status int, err error) error {
	return ctx.Status(status).JSON(fiber.Map{
		"message": err.Error(),
		"data":    nil,
	})
}

func InternalError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{
		"message": err.Error(),
		"data":    nil,
	})
}

func BadRequestError(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
		"message": msg,
		"data":    nil,
	})
}

func SuccessResponse(ctx *fiber.Ctx, msg string, data interface{}) error {
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": msg,
		"data":    data,
	})
}

func UnauthorizedError(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
		"message": msg,
		"data":    nil,
	})
}	

func NotFoundError(ctx *fiber.Ctx, msg string) error {
	return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
		"message": msg,	
		"data":    nil,
	})
}