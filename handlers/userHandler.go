package handlers

import (
	//"learning-backend/config"
	//"hash"
	"errors"
	dto_ "learning-backend/dto"
	"learning-backend/mapper"
	"learning-backend/rest"
	"learning-backend/service"
	"log"
	"net/http"
	"strconv"

	//"learning-backend/middleware"

	//"learning-backend/dto_"

	"github.com/gofiber/fiber/v2"
	//"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	//Service *service.UserService
		Service *service.UserService
}



func (h *UserHandler) SignUp(c *fiber.Ctx) error {
	var input dto_.SignUp //dto_.CreateUserdto_

	if err := c.BodyParser(&input); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}
	

	
	token, err := h.Service.SignUp(input)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "error on signup",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "register",
		"token":   token,
	})

}

func (h *UserHandler) SignIn(c *fiber.Ctx) error {
	var input dto_.SignIn

	if err := c.BodyParser(&input); err != nil {

		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	token, err := h.Service.SignIn(input.Email, input.Password)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "invalid credentials",
		})
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "login successful",
		"token":   token,
	})


}



func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}
	log.Println(user)
	// create verification code and update to user profile in DB
	error := h.Service.GetVerificationCode(user)
	log.Println(error)

	if error != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate verification code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
	})

}
func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}

	// request
	var req dto_.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	error := h.Service.VerifyCode(user.ID, req.Code)

	if error != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": error.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified successfully",
	})
}
func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}
	req := dto_.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}
	log.Printf("User %v", user)
	// create profile

	error := h.Service.CreateProfile(user.ID, req)

	if error != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to create profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile created successfully",
	})
}
func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}
	log.Println(user)

	// call user service and perform get profile
	profile, err := h.Service.GetProfile(user.ID)
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to get profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get profile",
		"profile": mapper.ToProfileResponse(profile),
	})
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}
	req := dto_.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	error := h.Service.UpdateProfile(user.ID, req)
	if error != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to update profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile updated successfully",
	})
}

func (h *UserHandler) AddtoCart(ctx *fiber.Ctx) error {

	req := dto_.CreateCartRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid product and qty",
		})
	}

	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}

	// call user service and perform create cart
	cartItems, err := h.Service.CreateCart(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "cart created successfully", mapper.ToCartResponseList(cartItems))

}
func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}
	cart, _, err := h.Service.FindCart(user.ID)
	if err != nil {
		return rest.InternalError(ctx, errors.New("cart does not exist"))
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
		"cart":    mapper.ToCartResponseList(cart),
	})
}

func (h *UserHandler) PlaceOrder(ctx *fiber.Ctx) error {

	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	req := dto_.PlaceOrderRequest{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "invalid request",
		})
	}

	err = h.Service.CreateOrder(
		user.ID,
		req.OrderRef,
		req.PaymentID,
	) // ❌ removed amount

	if err != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "order placed successfully",
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}

	orders, err := h.Service.GetOrders(user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
		"orders":  mapper.ToOrderResponseList(orders),
	})
}
func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	orderId, _ := strconv.Atoi(ctx.Params("id"))
	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}

	order, err := h.Service.GetOrderById(uint(orderId), user.ID)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order by id",
		"order":   mapper.ToOrderResponse(&order),
	})
}
func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	user, err := h.Service.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "unauthorized",
		})
	}

	req := dto_.SellerInput{}
	error := ctx.BodyParser(&req)
	if error != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "request parameters are not valid",
		})
	}

	token, err := h.Service.BecomeSeller(user.ID, req)


	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "fail to become seller",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become seller",
		"token":   token,
	})
}


func (h *UserHandler) GetShopBySlug(ctx *fiber.Ctx) error {
	slug := ctx.Params("shopSlug")

	shop, err := h.Service.GetShopBySlug(slug)
	if err != nil {
		return ctx.Status(http.StatusNotFound).JSON(&fiber.Map{
			"error": err.Error()})
	}

	return rest.SuccessResponse(ctx, "shop", mapper.ToShopPublicResponse(shop))
}



func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}