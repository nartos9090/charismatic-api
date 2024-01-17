package adapter

import (
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/productimage/entity"
)

type ProductImageRepoInterface interface {
	CreateProductImage(userID int, product CreateProductImageRepoReq) (int, *helpers_errors.Error)
	CreateGeneratedProductImage(userID int, generated AddGeneratedProductImageRepoReq) (*entity.ProductImageGenerated, *helpers_errors.Error)
	GetProductImageList(userID int) (*[]entity.ProductImage, *helpers_errors.Error)
	GetProductImage(userID, productID int) (*ProductImageDetail, *helpers_errors.Error)
	GetGeneratedProductImage(userID, id int) (*entity.ProductImageGenerated, *helpers_errors.Error)
}
