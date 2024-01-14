package entity

type Copywriting struct {
	ID           int    `json:"id" db:"id"`
	UserID       int    `json:"user_id" db:"user_id"`
	Title        string `json:"title" db:"title"`
	ProductImage string `json:"product_image" db:"product_image"`
	BrandName    string `json:"brand_name" db:"brand_name"`
	MarketTarget string `json:"market_target" db:"market_target"`
	Superiority  string `json:"superiority" db:"superiority"`
}
