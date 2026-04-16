package dto_

type SignUp struct {
	Name string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
	Phone string `json:"phone" validate:"required"`
}

type SignIn struct {
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`

}




type VerificationCodeInput struct {
	Code string `json:"code"`
}

type SellerInput struct {
	FirstName         string `json:"first_name"`
	LastName          string `json:"last_name"`
	ShopName          string `json:"shop_name"`
	PhoneNumber       string `json:"phone_number"`
	BankAccountNumber uint   `json:"bankAccountNumber"`
	GSTNumber         string `json:"gstNumber"`
	PANNumber         string `json:"panNumber"`
	SwiftCode         string `json:"swiftCode"`
	PaymentType       string `json:"paymentType"`
}

type AddressInput struct {
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	City         string `json:"city"`
	PostCode     uint   `json:"post_code"`
	Country      string `json:"country"`
}

type ProfileInput struct {
	FirstName    string       `json:"first_name"`
	LastName     string       `json:"last_name"`
	AddressInput AddressInput `json:"address"`
}