package adapter

import "go-api-echo/internal/services/video/entity"

type VideoProjectDetail struct {
	entity.VideoProject
	Scenes []entity.Scene `json:"scenes"`
}
