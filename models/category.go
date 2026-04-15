package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name         string    `json:"name" gorm:"index;"`
	ParentId     uint      `json:"parent_id"`
	ImageUrl     string    `json:"image_url" `
	Products     []Product `json:"products"`
	DisplayOrder int       `json:"display_order"`
}