package adapter

import "go-api-echo/internal/services/copywriting/entity"

type CreateCopywritingRepoReq struct {
	Title        string `json:"title" form:"title"`
	ProductImage string `json:"product_image" form:"product_image"`
	BrandName    string `json:"brand_name" form:"brand_name"`
	MarketTarget string `json:"market_target" form:"market_target"`
	Superiority  string `json:"superiority" form:"superiority"`
	Result       string `json:"result" form:"result"`
}

type CopywritingProjectDetail struct {
	entity.Copywriting
	Result *string `json:"result" db:"result"`
}
