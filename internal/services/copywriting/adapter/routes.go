package adapter

type CreateCopywritingReq struct {
	Title string `json:"title" form:"title"`
	//ProductImage multipart.File
	BrandName    string `json:"brand_name" form:"brand_name"`
	MarketTarget string `json:"market_target" form:"market_target"`
	Superiority  string `json:"superiority" form:"superiority"`
}
