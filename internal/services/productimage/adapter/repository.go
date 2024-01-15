package adapter

import "go-api-echo/internal/services/productimage/entity"

type CreateProductImageRepoReq struct {
	Title    string `json:"title" form:"title"`
	ImageUrl string
	MaskUrl  string
}

type AddGeneratedProductImageRepoReq struct {
	ProductImageID int    `json:"product_image_id" form:"product_image_id"`
	ImageUrl       string `json:"image_url" form:"image_url"`
	Prompt         string `json:"prompt" form:"prompt"`
}

type ProductImageDetail struct {
	entity.ProductImage
	GeneratedImages []entity.ProductImageGenerated `json:"generated_images"`
}
