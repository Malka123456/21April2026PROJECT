package models

import "gorm.io/gorm"

type Shop struct {
	gorm.Model

	UserID uint `json:"user_id"`
	User   User `gorm:"foreignKey:UserID"` // 👈 important

	Name string `json:"name"`
	Slug string `json:"slug" gorm:"unique"`

	GSTNumber string `json:"gst_number"`
	PANNumber string `json:"pan_number"`

	Description string `json:"description"`
	LogoURL     string `json:"logo_url"`

	IsVerified bool `json:"is_verified" gorm:"default:false"`
	IsActive   bool `json:"is_active" gorm:"default:true"`

	// relation
	Products []Product `json:"products"`
}
