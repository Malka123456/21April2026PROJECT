package dto_

type ShopResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type ProductResponse struct {
	ID    uint        `json:"id"`
	Name  string      `json:"name"`
	Price float64     `json:"price"`
	Shop  ShopResponse `json:"shop"`
}