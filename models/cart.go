package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserId    uint      `json:"user_id"`
	User      *User      `gorm:"foreignKey:UserID"`
	ProductId uint      `json:"product_id"`
	Name      string    `json:"name"`
	ImageUrl  string    `json:"image_url"`
	SellerId  uint      `json:"seller_id"`
	Price     float64   `json:"price"`
	Qty       uint      `json:"qty"`
}