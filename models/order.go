package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserId         uint        `json:"user_id"`
	User           User        `gorm:"foreignKey:UserID"`
	Status         string      `json:"status"`
	Amount         float64     `json:"amount"`
	TransactionId  string      `json:"transaction_id"`
	OrderRefNumber string      `json:"order_ref_number"`
	PaymentId      string      `json:"payment_id"`
	Items          []OrderItem `json:"items"`
}
