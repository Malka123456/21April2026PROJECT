package dto_

type CategoryResponse struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
	ParentID uint   `json:"parent_id"`
}