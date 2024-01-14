package adapter

import (
	"fmt"
	"go-api-echo/external/services/gemini_service"
	"go-api-echo/internal/pkg/helpers/response"
	"net/http"
)

func CreateCopywritingProjectSync(userID int, req CreateCopywritingReq, repo CopywritingProjectRepoInterface) (resp response.HttpRes) {
	result, err := gemini_service.GenerateCopywriting(gemini_service.GenerateCopywritingRequest{
		ProductImage: req.ProductImage,
		BrandName:    req.BrandName,
		MarketTarget: req.MarketTarget,
		Superiority:  req.Superiority,
	})
	if err != nil {
		return err.ToHttpRes()
	}

	fmt.Println(result)

	project, err := repo.Create(userID, CreateCopywritingRepoReq{
		Title:        req.Title,
		BrandName:    req.BrandName,
		MarketTarget: req.MarketTarget,
		Superiority:  req.Superiority,
		Result:       result,
	})
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Status = http.StatusOK
	resp.Message = "data created successfully"
	resp.Data = project
	return resp
}

func GetCopywritingList(userID int, repo CopywritingProjectRepoInterface) (resp response.HttpResPaginated) {
	projects, err := repo.GetLists(userID)
	if err != nil {
		return response.HttpResPaginated(err.ToHttpRes())
	}

	resp.Status = http.StatusOK
	resp.Message = "data retrieved successfully"
	resp.Data = projects
	return resp
}

func GetCopywritingDetail(userID int, projectID int, repo CopywritingProjectRepoInterface) (resp response.HttpRes) {
	project, err := repo.GetDetail(projectID, userID)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Status = http.StatusOK
	resp.Message = "data retrieved successfully"
	resp.Data = project
	return resp
}
