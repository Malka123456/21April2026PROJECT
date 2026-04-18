package dto_

type CartResponse struct {
	ID        uint    `json:"id"`
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	ImageURL  string  `json:"image_url"`
	Price     float64 `json:"price"`
	Qty       uint     `json:"qty"`
}

type OrderResponse struct {
	ID        uint    `json:"id"`
	UserID    uint    `json:"user_id"`
	Total     float64 `json:"total"`
	Items []OrderItemResponse `json:"items"`
	Status    string  `json:"status"`
	CreatedAt string  `json:"created_at"`
}

type OrderItemResponse struct {
	ProductID uint    `json:"product_id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Qty       uint     `json:"qty"`
}