package adapter

import (
	"go-api-echo/internal/pkg/helpers/response"
)

func GetHistory(userID int, repo HistoryRepoInterface) (resp response.HttpRes) {
	histories, err := repo.GetHistory(userID)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Message = "Success"
	resp.Data = histories
	return resp
}
