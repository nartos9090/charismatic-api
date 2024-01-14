package adapter

import (
	"go-api-echo/external/services/dalle_service"
	"go-api-echo/external/services/elevenlabs_service"
	"go-api-echo/external/services/gemini_service"
	"go-api-echo/internal/pkg/helpers/response"
	"go-api-echo/internal/services/video/entity"
	"log"
	"net/http"
)

func AddVideoProject(userID int, req GenerateVideoReq, repo VideoProjectRepoInterface) (resp response.HttpRes) {
	project, err := repo.CreateVideoProject(userID, &req)
	if err != nil {
		return err.ToHttpRes()
	}

	go GenerateVideo(project.ID, *project, repo)

	resp.Status = http.StatusOK
	resp.Message = "data created successfully"
	resp.Data = project

	return
}

func GenerateVideo(projectID int, project entity.VideoProject, repo VideoProjectRepoInterface) {
	storyboards, err := gemini_service.Generate(gemini_service.GenerateRequest{
		ProductTitle: project.ProductTitle,
		BrandName:    project.BrandName,
		ProductType:  project.ProductType,
		MarketTarget: project.MarketTarget,
		Superiority:  project.Superiority,
		Duration:     project.Duration,
	})
	if err != nil {
		log.Println("error generating storyboard")
		log.Print(err.Message)
		return
	}

	for i, storyboard := range storyboards {
		go func(i int, storyboard gemini_service.Storyboard) {
			sceneID, err := repo.CreateScene(projectID, &entity.Scene{
				Sequence:        i + 1,
				Title:           storyboard.Title,
				Narration:       storyboard.Narration,
				Illustration:    storyboard.Illustration,
				IllustrationUrl: storyboard.IllustrationUrl,
				VoiceUrl:        storyboard.VoiceUrl,
			})

			if err != nil {
				log.Println("error create scene (projectID: ", projectID, ")")
				log.Print(err.Message)
				return
			}

			log.Println("created scene (sceneID: ", sceneID, ")")

			go func(sceneID int, storyBoard gemini_service.Storyboard) {
				url, err := dalle_service.Generate(storyBoard.Illustration, dalle_service.GenerateSize1)
				if err != nil {
					log.Print(err.Message)
					return
				}

				log.Println("generated illustration (sceneID: ", sceneID, ")")

				repo.UpdateScene(sceneID, &entity.Scene{
					IllustrationUrl: url,
				})
			}(sceneID, storyboard)

			go func(sceneID int, storyBoard gemini_service.Storyboard) {
				url, err := elevenlabs_service.Generate(storyBoard.Narration)
				if err != nil {
					log.Print(err.Message)
					return
				}

				log.Println("generated illustration (sceneID: ", sceneID, ")")

				repo.UpdateScene(sceneID, &entity.Scene{
					VoiceUrl: url,
				})
			}(sceneID, storyboard)
		}(i, storyboard)
	}
}
