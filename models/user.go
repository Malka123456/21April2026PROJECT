package models

import (
	"time"

	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin  Role = "admin"
	RoleSeller Role = "seller"
	RoleBuyer  Role = "buyer"
)

type User struct {
	gorm.Model
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email" gorm:"index;unique;not null"`
	Phone     string    `json:"phone"`
	Password  string    `json:"password"`
	Code      string    `json:"code"`
	Expiry    time.Time `json:"expiry"`
	Address   Address   `json:"address"` // relation
	Cart      Cart      `json:"cart"`    // relation

	Orders   []Order   `json:"orders"`   // relation
	Payments []Payment `json:"payment" ` // relation
	Verified bool      `json:"verified" gorm:"default:false"`
	UserType Role      `json:"user_type" gorm:"default:buyer"`
}
