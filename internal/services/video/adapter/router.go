package adapter

type (
	GenerateVideoReq struct {
		ProductTitle string `json:"product_title"`
		BrandName    string `json:"brand_name"`
		ProductType  string `json:"product_type"`
		MarketTarget string `json:"market_target"`
		Superiority  string `json:"superiority"`
		Duration     int    `json:"duration"`
	}
)
