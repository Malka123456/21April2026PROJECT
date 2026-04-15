package models

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	AddressLine1 string    `json:"address_line1"`
	AddressLine2 string    `json:"address_line2"`
	City         string    `json:"city"`
	PostCode     uint      `json:"postCode"`
	Country      string    `json:"country"`
	UserId       uint      `json:"user_id"`
}