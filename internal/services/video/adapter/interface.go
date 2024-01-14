package adapter

import (
	errors "go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/video/entity"
)

type VideoProjectRepoInterface interface {
	CreateVideoProject(userID int, req *GenerateVideoReq) (*entity.VideoProject, *errors.Error)
	GetVideoProjectList(userID int) (*[]entity.VideoProject, *errors.Error)
	GetVideoProjectDetail(projectID int, userID int) (*VideoProjectDetail, *errors.Error)

	CreateScene(projectID int, scene *entity.Scene) (int, *errors.Error)
	UpdateScene(sceneID int, scene *entity.Scene) (int, *errors.Error)
}
