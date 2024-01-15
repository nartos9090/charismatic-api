package adapter

type CreateProductImageReq struct {
	Prompt    string `json:"prompt" form:"prompt"`
	Title     string `json:"title" form:"title"`
	BaseImage string
	MaskImage string
}

type GenerateBackgroundReq struct {
	ProductImageID int    `param:"id" json:"product_image_id" form:"product_image_id"`
	Prompt         string `json:"prompt" form:"prompt"`
}
