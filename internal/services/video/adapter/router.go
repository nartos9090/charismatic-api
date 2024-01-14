package adapter

type (
	GenerateVideoReq struct {
		ProductTitle string `json:"product_title" form:"product_title" validate:"required"`
		BrandName    string `json:"brand_name" form:"brand_name" validate:"required"`
		ProductType  string `json:"product_type" form:"product_type" validate:"required"`
		MarketTarget string `json:"market_target" form:"market_target" validate:"required"`
		Superiority  string `json:"superiority" form:"superiority" validate:"required"`
		Duration     int    `json:"duration" form:"duration" validate:"required"`
	}
)
