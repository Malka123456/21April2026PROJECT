package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string    `json:"name" gorm:"index;"`
	Description string    `json:"description"`
	CategoryId  uint      `json:"category_id"`
	ImageUrl    string    `json:"image_url" `
	Price       float64   `json:"price"`
	UserId      int       `json:"user_id"`
	Stock       uint      `json:"stock"`
}