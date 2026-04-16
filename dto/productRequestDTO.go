package dto_

type CreateProductRequest struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	CategoryID  uint    `json:"category_id"`
	ImageURL    string  `json:"image_url"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
}

type UpdateStockRequest struct {
	Stock int `json:"stock"`
}