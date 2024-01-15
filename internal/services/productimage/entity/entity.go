package entity

type ProductImage struct {
	ID       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id" db:"user_id"`
	Title    string `json:"title" db:"title"`
	ImageUrl string `json:"image_url" db:"image_url"`
	MaskUrl  string `json:"mask_url" db:"mask_url"`
}

type ProductImageGenerated struct {
	ID             int    `json:"id" db:"id"`
	ProductImageID int    `json:"product_image_id" db:"product_image_id"`
	ImageUrl       string `json:"image_url" db:"image_url"`
	Prompt         string `json:"prompt" db:"prompt"`
}
