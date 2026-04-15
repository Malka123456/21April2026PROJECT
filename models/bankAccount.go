package models

import "gorm.io/gorm"

type BankAccount struct {
	gorm.Model
	UserId      uint      `json:"user_id"`
	BankAccount uint      `json:"bank_account" gorm:"index;unique;not null"`
	SwiftCode   string    `json:"swift_code"`
	PaymentType string    `json:"payment_type"`
}