package mapper

import (
	dto_ "learning-backend/dto"
	"learning-backend/models"
)

//
// ==================== PRODUCT (SELLER) ====================
//

func ToProductSellerResponse(p *models.Product) dto_.ProductSellerResponse {
	return dto_.ProductSellerResponse{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
		Slug:  p.Slug,
	}
}

func ToProductSellerResponseList(products []*models.Product) []dto_.ProductSellerResponse {
	result := make([]dto_.ProductSellerResponse, len(products))
	for i := range products {
		result[i] = ToProductSellerResponse(products[i])
	}
	return result
}

//
// ==================== PRODUCT (PUBLIC) ====================
//

func ToProductPublicResponse(p *models.Product) dto_.ProductPublicResponse {
	return dto_.ProductPublicResponse{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
		Image: p.ImageURL,
	}
}

func ToProductPublicResponseList(products []*models.Product) []dto_.ProductPublicResponse {
	result := make([]dto_.ProductPublicResponse, len(products))
	for i := range products {
		result[i] = ToProductPublicResponse(products[i])
	}
	return result
}

//
// ==================== SHOP (SELLER) ====================
//

func ToShopSellerResponse(s models.Shop) dto_.ShopSellerResponse {
	return dto_.ShopSellerResponse{
		ID:       s.ID,
		Name:     s.Name,
		Slug:     s.Slug,
		Products: ToProductSellerResponseList(toConvertInPointer(s.Products)),
	}
}

//
// ==================== SHOP (PUBLIC) ====================
//

func ToShopPublicResponse(s models.Shop) dto_.ShopPublicResponse {
	return dto_.ShopPublicResponse{
		ID:       s.ID,
		Name:     s.Name,
		Products: ToProductPublicResponseList(toConvertInPointer(s.Products)),
	}
}

func toConvertInPointer(prod []models.Product) []*models.Product {
	var ptrs []*models.Product
for i := range prod {
	ptrs = append(ptrs, &prod[i])
}
return ptrs
}




func ToCategoryResponse(c *models.Category) dto_.CategoryResponse {
	return dto_.CategoryResponse{
		ID:       c.ID,
		Name:     c.Name,
		ImageURL: c.ImageURL,
		ParentID: c.ParentID,
	}
}

func ToCategoryResponseList(categories []*models.Category) []dto_.CategoryResponse {
	result := make([]dto_.CategoryResponse, len(categories))
	for i := range categories {
		result[i] = ToCategoryResponse(categories[i])
	}
	return result
}

func ToProfileResponse(user *models.User) *dto_.ProfileResponse {
	return &dto_.ProfileResponse{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.Phone,
		UserType:  string(user.UserType),
		Address:   ToAddressResponse(&user.Address),
	}
}




func ToAddressResponse(a *models.Address) dto_.AddressResponse {
	return dto_.AddressResponse{
		AddressLine1: a.AddressLine1,
		AddressLine2: a.AddressLine2,
		City:         a.City,
		PostCode:     a.PostCode,
		Country:      a.Country,
	}
}			

func ToCartResponse(c *models.Cart) dto_.CartResponse {
	return dto_.CartResponse{
		ID:        c.ID,
		ProductID: c.ProductID,
		Name:      c.Name,
		ImageURL:  c.ImageURL,
		Price:     c.Price,
		Qty:       c.Qty,
	}
}

func ToCartResponseList(cartItems []models.Cart) []dto_.CartResponse {
	result := make([]dto_.CartResponse, len(cartItems))
	for i := range cartItems {
		result[i] = ToCartResponse(&cartItems[i])
	}
	return result
}


func ToOrderResponse(o *models.Order) dto_.OrderResponse {
	return dto_.OrderResponse{
		ID:        o.ID,
		UserID:    o.UserID,
		Total:     o.Amount,
		Items:     ToOrderItemResponseList(o.Items),
		Status:    string(o.Status),
		CreatedAt: o.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func ToOrderResponseList(orders []models.Order) []dto_.OrderResponse {
	result := make([]dto_.OrderResponse, len(orders))
	for i := range orders {
		result[i] = ToOrderResponse(&orders[i])
	}
	return result
}

func ToOrderItemResponse(i *models.OrderItem) dto_.OrderItemResponse {
	return dto_.OrderItemResponse{
		ProductID: i.ProductID,
		Name:      i.Name,
		Price:     i.Price,
		Qty:       i.Qty,
	}
}

func ToOrderItemResponseList(items []models.OrderItem) []dto_.OrderItemResponse {
	result := make([]dto_.OrderItemResponse, len(items))
	for i := range items {
		result[i] = ToOrderItemResponse(&items[i])
	}
	return result
}