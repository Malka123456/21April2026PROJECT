package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID    uint      `json:"user_id"`
	User      *User      `gorm:"foreignKey:UserID"`
	ProductID uint      `json:"product_id"`
	Name      string    `json:"name"`
	ImageURL  string    `json:"image_url"`
	SellerID  uint      `json:"seller_id"`
	Price     float64   `json:"price"`
	Qty       uint      `json:"qty"`
}