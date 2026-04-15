package handlers

import (
	//"learning-backend/config"
	//"hash"
	"errors"
	dto_ "learning-backend/dto"
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

// func (h *UserHandler) GetVerificationCode(c *fiber.Ctx) error {
// 	return nil
// }

// func (h *UserHandler) Verify(c *fiber.Ctx) error {
// 	return nil
// }

// func (h *UserHandler) CreateProfile(c *fiber.Ctx) error {
// 	return nil
// }


// func (h *UserHandler) GetProfile(c *fiber.Ctx) error {

// 	return nil
// }

// func (h *UserHandler) UpdateProfile(c *fiber.Ctx) error {

// 	return nil
// }

// func (h *UserHandler) Addto_Cart(c *fiber.Ctx) error {
// 	return nil
// }

// func (h *UserHandler) GetCart(c *fiber.Ctx) error {
// 	return nil
// }

// func (h *UserHandler) GetOrders(c *fiber.Ctx) error {
// 	return nil
// }

// func (h *UserHandler) GetOrder(c *fiber.Ctx) error {
// 	return nil
// }

// func (h *UserHandler) BecomeSeller(c *fiber.Ctx) error {
// 	return nil
// }

func (h *UserHandler) GetVerificationCode(ctx *fiber.Ctx) error {

	user := h.Service.Auth.GetCurrentUser(ctx)
	log.Println(user)
	// create verification code and update to user profile in DB
	err := h.Service.GetVerificationCode(user)
	log.Println(err)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to generate verification code",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get verification code",
	})

}
func (h *UserHandler) Verify(ctx *fiber.Ctx) error {

	user := h.Service.Auth.GetCurrentUser(ctx)

	// request
	var req dto_.VerificationCodeInput

	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err := h.Service.VerifyCode(user.ID, req.Code)

	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "verified successfully",
	})
}
func (h *UserHandler) CreateProfile(ctx *fiber.Ctx) error {

	user := h.Service.Auth.GetCurrentUser(ctx)
	req := dto_.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}
	log.Printf("User %v", user)
	// create profile

	err := h.Service.CreateProfile(user.ID, req)

	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(&fiber.Map{
			"message": "unable to create profile",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "profile created successfully",
	})
}
func (h *UserHandler) GetProfile(ctx *fiber.Ctx) error {

	user := h.Service.Auth.GetCurrentUser(ctx)
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
		"profile": profile,
	})
}

func (h *UserHandler) UpdateProfile(ctx *fiber.Ctx) error {
	user := h.Service.Auth.GetCurrentUser(ctx)
	req := dto_.ProfileInput{}
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(&fiber.Map{
			"message": "please provide a valid input",
		})
	}

	err := h.Service.UpdateProfile(user.ID, req)
	if err != nil {
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

	user := h.Service.Auth.GetCurrentUser(ctx)

	// call user service and perform create cart
	cartItems, err := h.Service.CreateCart(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "cart created successfully", cartItems)

}
func (h *UserHandler) GetCart(ctx *fiber.Ctx) error {
	user := h.Service.Auth.GetCurrentUser(ctx)
	cart, _, err := h.Service.FindCart(user.ID)
	if err != nil {
		return rest.InternalError(ctx, errors.New("cart does not exist"))
	}
	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get cart",
		"cart":    cart,
	})
}

func (h *UserHandler) GetOrders(ctx *fiber.Ctx) error {
	user := h.Service.Auth.GetCurrentUser(ctx)

	orders, err := h.Service.GetOrders(user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get orders",
		"orders":  orders,
	})
}
func (h *UserHandler) GetOrder(ctx *fiber.Ctx) error {
	orderId, _ := strconv.Atoi(ctx.Params("id"))
	user := h.Service.Auth.GetCurrentUser(ctx)

	order, err := h.Service.GetOrderById(uint(orderId), user.ID)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "get order by id",
		"order":   order,
	})
}
func (h *UserHandler) BecomeSeller(ctx *fiber.Ctx) error {

	user := h.Service.Auth.GetCurrentUser(ctx)

	req := dto_.SellerInput{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(&fiber.Map{
			"message": "request parameters are not valid",
		})
	}

	token, err := h.Service.BecomeSeller(user.ID, req)

	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
			"message": "fail to become seller",
		})
	}

	return ctx.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "become seller",
		"token":   token,
	})
}





func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}