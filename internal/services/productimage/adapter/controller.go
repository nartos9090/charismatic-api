package adapter

import (
	"fmt"
	"go-api-echo/config"
	"go-api-echo/external/services/background_removal"
	"go-api-echo/external/services/background_replacer"
	"go-api-echo/internal/pkg/helpers/response"
)

func CreateProductImage(userID int, req CreateProductImageReq, repo ProductImageRepoInterface, backgroundReplacerRepo background_replacer.BackgroundReplacerRepoInterface, backgroundRemoverRepo background_removal.BackgroundRemovalRepoInterface) (resp response.HttpRes) {
	fmt.Println(":: CreateProductImage")
	removedBackgroundImageUrl, err := backgroundRemoverRepo.RemoveBackground(req.BaseImage, req.MaskImage, req.Prompt)
	if err != nil {
		return err.ToHttpRes()
	}

	fmt.Println(":: removedBackgroundImageUrl", removedBackgroundImageUrl)
	productID, err := repo.CreateProductImage(userID, CreateProductImageRepoReq{
		Title:    req.Title,
		ImageUrl: removedBackgroundImageUrl,
		MaskUrl:  req.MaskImage,
	})

	if err != nil {
		return err.ToHttpRes()
	}

	return GenerateImageBackground(
		userID,
		GenerateBackgroundReq{
			Prompt:         req.Prompt,
			ProductImageID: productID,
		},
		repo,
		backgroundReplacerRepo,
	)
}

func GenerateImageBackground(userID int, req GenerateBackgroundReq, repo ProductImageRepoInterface, backgroundReplacerRepo background_replacer.BackgroundReplacerRepoInterface) (resp response.HttpRes) {
	product, err := repo.GetProductImage(userID, req.ProductImageID)
	if err != nil {
		return err.ToHttpRes()
	}

	imageUrl, err := backgroundReplacerRepo.ReplaceBackground(product.ImageUrl, product.MaskUrl, req.Prompt)
	if err != nil {
		return err.ToHttpRes()
	}

	generatedImage, err := repo.CreateGeneratedProductImage(userID, AddGeneratedProductImageRepoReq{
		ImageUrl:       imageUrl,
		Prompt:         req.Prompt,
		ProductImageID: product.ID,
	})
	if err != nil {
		return err.ToHttpRes()
	}

	generatedImage.ImageUrl = config.GlobalEnv.BaseURL + generatedImage.ImageUrl

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

	product.ImageUrl = config.GlobalEnv.BaseURL + product.ImageUrl
	product.MaskUrl = config.GlobalEnv.BaseURL + product.MaskUrl

	for i := 0; i < len(product.GeneratedImages); i++ {
		product.GeneratedImages[i].ImageUrl = config.GlobalEnv.BaseURL + product.GeneratedImages[i].ImageUrl
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

	for i := 0; i < len(*products); i++ {
		(*products)[i].ImageUrl = config.GlobalEnv.BaseURL + (*products)[i].ImageUrl
		(*products)[i].MaskUrl = config.GlobalEnv.BaseURL + (*products)[i].MaskUrl
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

	generatedImage.ImageUrl = config.GlobalEnv.BaseURL + generatedImage.ImageUrl

	return response.HttpRes{
		Status:  200,
		Message: "Success",
		Data:    generatedImage,
		Errors:  nil,
	}
}
