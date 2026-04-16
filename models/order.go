package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID         uint        `json:"user_id"`
	User           User        `gorm:"foreignKey:UserID"`
	Status         string      `json:"status"`
	Amount         float64     `json:"amount"`
	TransactionID  string      `json:"transaction_id"`
	OrderRefNumber string      `json:"order_ref_number"`
	PaymentID      string      `json:"payment_id"`
	Items          []OrderItem `json:"items"`
}
