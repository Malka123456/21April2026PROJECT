package dto_


type UserResponseDTO struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

type AddressResponse struct {
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	PostCode     uint   `json:"post_code"`
	Country      string `json:"country"`
}

type ProfileResponse struct {
	ID        uint            `json:"id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Email     string          `json:"email"`
	Phone     string          `json:"phone"`
	UserType  string          `json:"user_type"`
	Address   AddressResponse `json:"address"`
}


