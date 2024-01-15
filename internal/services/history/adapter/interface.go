package adapter

import (
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/history/entity"
)

type HistoryRepoInterface interface {
	GetHistory(userID int) (*[]entity.History, *helpers_errors.Error)
}
