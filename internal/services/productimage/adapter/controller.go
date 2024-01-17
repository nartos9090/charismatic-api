package adapter

import (
	"go-api-echo/external/services/dalle_service"
	"go-api-echo/internal/pkg/helpers/response"
)

func CreateProductImage(userID int, req CreateProductImageReq, repo ProductImageRepoInterface) (resp response.HttpRes) {
	productID, err := repo.CreateProductImage(userID, CreateProductImageRepoReq{
		Title:    req.Title,
		ImageUrl: req.BaseImage,
		MaskUrl:  req.MaskImage,
	})
	if err != nil {
		return err.ToHttpRes()
	}

	return GenerateImageBackground(userID, GenerateBackgroundReq{
		Prompt:         req.Prompt,
		ProductImageID: productID,
	}, repo)
}

func GenerateImageBackground(userID int, req GenerateBackgroundReq, repo ProductImageRepoInterface) (resp response.HttpRes) {
	product, err := repo.GetProductImage(userID, req.ProductImageID)
	if err != nil {
		return err.ToHttpRes()
	}

	imageUrl, err := dalle_service.GenerateBackground(product.ImageUrl, product.MaskUrl, req.Prompt)
	if err != nil {
		return err.ToHttpRes()
	}

	generatedImage, err := repo.CreateGeneratedProductImage(product.ID, AddGeneratedProductImageRepoReq{
		ImageUrl:       imageUrl,
		Prompt:         req.Prompt,
		ProductImageID: product.ID,
	})
	if err != nil {
		return err.ToHttpRes()
	}

	return response.HttpRes{
		Status:  200,
		Message: "Success",
		Data:    generatedImage,
		Errors:  nil,
	}
}

func GetProductImage(userID, productID int, repo ProductImageRepoInterface) (resp response.HttpRes) {
	product, err := repo.GetProductImage(userID, productID)
	if err != nil {
		return err.ToHttpRes()
	}

	return response.HttpRes{
		Status:  200,
		Message: "Success",
		Data:    product,
		Errors:  nil,
	}
}

func GetProductImageList(userID int, repo ProductImageRepoInterface) (resp response.HttpRes) {
	products, err := repo.GetProductImageList(userID)
	if err != nil {
		return err.ToHttpRes()
	}

	return response.HttpRes{
		Status:  200,
		Message: "Success",
		Data:    products,
		Errors:  nil,
	}
}

func GetGeneratedProductImage(userID, generatedImageID int, repo ProductImageRepoInterface) (resp response.HttpRes) {
	generatedImage, err := repo.GetGeneratedProductImage(userID, generatedImageID)
	if err != nil {
		return err.ToHttpRes()
	}

	return response.HttpRes{
		Status:  200,
		Message: "Success",
		Data:    generatedImage,
		Errors:  nil,
	}
}
