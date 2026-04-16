package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID   uint      `json:"order_id"`
	UserID    uint      `json:"user_id"`     
	ProductID uint      `json:"product_id"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	SellerID  uint      `json:"seller_id"`
	Price     float64   `json:"price"`
	Qty       uint      `json:"qty"`
}