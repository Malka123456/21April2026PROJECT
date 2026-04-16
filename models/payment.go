package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID        uint          `json:"user_id"`
	CaptureMethod string        `json:"capture_method"`
	Amount        float64       `json:"amount"`
	OrderID       string        `json:"order_id"`
	CustomerID    string        `json:"customer_id"` // stripe customer if
	PaymentID     string        `json:"payment_id"`  // payment id
	ClientSecret  string        `json:"client_secret"`
	Status        PaymentStatus `json:"status" gorm:"default:initial"` // initial, success, failed
	Response      string        `json:"response"`
}

type PaymentStatus string 

const (
	PaymentStatusInitial PaymentStatus = "initial"
	PaymentStatusSuccess PaymentStatus = "success"
	PaymentStatusFailed  PaymentStatus = "failed"
	PaymentStatusPending PaymentStatus = "pending"
)