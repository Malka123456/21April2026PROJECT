package handlers

import (
	"errors"
	dto_ "learning-backend/dto"
	"learning-backend/mapper"
	"learning-backend/models"
	"learning-backend/rest"
	"learning-backend/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type CatalogHandler struct {
	svc *service.CatalogService
}

func (h CatalogHandler) GetCategories(ctx *fiber.Ctx) error {

	cats, err := h.svc.GetCategories()
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}
	return rest.SuccessResponse(ctx, "categories", mapper.ToCategoryResponseList(cats))
}
func (h CatalogHandler) GetCategoryById(ctx *fiber.Ctx) error {

id, err := strconv.Atoi(ctx.Params("id"))
if err != nil {
	return rest.BadRequestError(ctx, "invalid id")
}
	cat, err := h.svc.GetCategory(id)
	if err != nil {
		return rest.ErrorMessage(ctx, 404, err)
	}
	return rest.SuccessResponse(ctx, "category", mapper.ToCategoryResponse(cat))
}

func (h CatalogHandler) CreateCategories(ctx *fiber.Ctx) error {

	req := dto_.CreateCategoryRequest{}

	err := ctx.BodyParser(&req)

	if err != nil {
		return rest.BadRequestError(ctx, "create category request is not valid")
	}

	err = h.svc.CreateCategory(req)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "category created successfully", nil)
}

func (h CatalogHandler) EditCategory(ctx *fiber.Ctx) error {

id, err := strconv.Atoi(ctx.Params("id"))
if err != nil {
	return rest.BadRequestError(ctx, "invalid id")
}
	req := dto_.CreateCategoryRequest{}

	error := ctx.BodyParser(&req)

	if error != nil {
		return rest.BadRequestError(ctx, "update category request is not valid")
	}

	updatedCat, err := h.svc.EditCategory(id, req)

	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "edit category", mapper.ToCategoryResponse(updatedCat))
}

func (h CatalogHandler) DeleteCategory(ctx *fiber.Ctx) error {
id, err := strconv.Atoi(ctx.Params("id"))
if err != nil {
	return rest.BadRequestError(ctx, "invalid id")
}
	error := h.svc.DeleteCategory(id)
	if error != nil {
		return rest.InternalError(ctx, error)
	}
	return rest.SuccessResponse(ctx, "category deleted successfully", nil)
}

func (h CatalogHandler) CreateProducts(ctx *fiber.Ctx) error {

	req := dto_.CreateProductRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return rest.BadRequestError(ctx, "create product request is not valid")
	}

	user, err := h.svc.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(401).JSON(&fiber.Map{
			"message": err.Error(),
		})
	}

	product, err := h.svc.CreateProduct(req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "product created successfully", mapper.ToProductSellerResponse(product))
}

func (h CatalogHandler) GetProductsForPublic(ctx *fiber.Ctx) error {

	products, err := h.svc.GetProducts()
	if err != nil {
		return rest.BadRequestError(ctx, "products not found")
	}

	return rest.SuccessResponse(ctx, "products", mapper.ToProductPublicResponseList(products))
}

func (h CatalogHandler) GetProductsForSeller(ctx *fiber.Ctx) error {

	// get logged-in user
	user, err := h.svc.Auth.GetCurrentUser(ctx)
	if err != nil {
		return rest.UnauthorizedError(ctx, "unauthorized")
	}

	products, err := h.svc.GetSellerProducts(int(user.ID))
	if err != nil {
		return rest.BadRequestError(ctx, "products not found")
	}

	return rest.SuccessResponse(ctx, "products", mapper.ToProductSellerResponseList(products))
}


func (h CatalogHandler) GetProductForPublic(ctx *fiber.Ctx) error {

id, err := strconv.Atoi(ctx.Params("id"))
if err != nil {
	return rest.BadRequestError(ctx, "invalid id")
}
	product, err := h.svc.GetProductById(id)
	if err != nil {
		
		if errors.Is(err, service.ErrProductNotFound) {
			return rest.NotFoundError(ctx, "product not found")
		}
		
		return rest.InternalError(ctx, err)
	}

	return rest.SuccessResponse(ctx, "product", mapper.ToProductPublicResponse(product))
}

func (h CatalogHandler) GetProductForSeller(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return rest.BadRequestError(ctx, "invalid id")
	}
	user, err := h.svc.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}


	// get product
	product, err := h.svc.GetProductById(id)
	if err != nil {
		return rest.BadRequestError(ctx, "product not found")
	}

	// 🔥 IMPORTANT: check ownership
	if product.Shop.UserID != user.ID {
		return rest.BadRequestError(ctx, "you are not allowed to access this product")
	}

	return rest.SuccessResponse(ctx, "product",  mapper.ToProductSellerResponse(product))
}
// GetProductBySlug gets product by slug
func (h CatalogHandler) GetProductBySlug(ctx *fiber.Ctx) error {
	shopSlug := ctx.Params("shopSlug")
	productSlug := ctx.Params("productSlug")

	product, err := h.svc.GetProductBySlug(shopSlug, productSlug)
	if err != nil {
		return rest.BadRequestError(ctx, "product not found")
	}

	return rest.SuccessResponse(ctx, "product", mapper.ToProductPublicResponse(product))
}

func (h CatalogHandler) EditProduct(ctx *fiber.Ctx) error {

id, err := strconv.Atoi(ctx.Params("id"))
if err != nil {
	return rest.BadRequestError(ctx, "invalid id")
}	

req := dto_.CreateProductRequest{}
	error := ctx.BodyParser(&req)
	if error != nil {
		return rest.BadRequestError(ctx, "edit product request is not valid")
	}
	user, err := h.svc.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	product, err := h.svc.EditProduct(id, req, user)
	if err != nil {
		return rest.InternalError(ctx, err)
	}
	return rest.SuccessResponse(ctx, "edit product", mapper.ToProductSellerResponse(product))
}

func (h CatalogHandler) UpdateStock(ctx *fiber.Ctx) error {
id, err := strconv.Atoi(ctx.Params("id"))
if err != nil {
	return rest.BadRequestError(ctx, "invalid id")
}

req := dto_.UpdateStockRequest{}
	error := ctx.BodyParser(&req)
	if error != nil {
		return rest.BadRequestError(ctx, "update stock request is not valid")
	}
	user, err := h.svc.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	product := models.Product{
		CategoryID: uint(id),
		Stock:      uint(req.Stock),
		ShopID:     uint(user.ID),
	}

	updatedProduct, err := h.svc.UpdateProductStock(product)

	return rest.SuccessResponse(ctx, "update stock ", mapper.ToProductSellerResponse(updatedProduct))
}

func (h CatalogHandler) DeleteProduct(ctx *fiber.Ctx) error {

id, err := strconv.Atoi(ctx.Params("id"))
if err != nil {
	return rest.BadRequestError(ctx, "invalid id")
}	// need to provide user id to verify ownership
	user, err := h.svc.Auth.GetCurrentUser(ctx)
	if err != nil {
		return ctx.Status(401).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	error := h.svc.DeleteProduct(id, user)

	return rest.SuccessResponse(ctx, "Delete product ", error)
}

func NewCatalogHandler(svc *service.CatalogService) *CatalogHandler {
	return &CatalogHandler{
		svc: svc,
	}
}
