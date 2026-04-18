package dto_


// Public DTO (normal user)
type ProductPublicResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type ShopPublicResponse struct {
	ID       uint                    `json:"id"`
	Name     string                  `json:"name"`
	Products []ProductPublicResponse `json:"products"`
}


// Seller DTO (for seller)
type ProductSellerResponse struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Slug  string  `json:"slug"`
}

type ShopSellerResponse struct {
	ID       uint                     `json:"id"`
	Name     string                   `json:"name"`
	Slug     string                   `json:"slug"`
	Products []ProductSellerResponse  `json:"products"`
}