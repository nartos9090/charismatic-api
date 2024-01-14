package adapter

import (
	errors "go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/copywriting/entity"
)

type CopywritingProjectRepoInterface interface {
	Create(userID int, project CreateCopywritingRepoReq) (*CopywritingProjectDetail, *errors.Error)
	GetLists(userID int) (*[]entity.Copywriting, *errors.Error)
	GetDetail(projectID int, userID int) (*CopywritingProjectDetail, *errors.Error)
}
