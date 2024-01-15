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

func AddVideoProjectSyncWithDetail(userID int, req GenerateVideoReq, repo VideoProjectRepoInterface) (resp response.HttpRes) {
	project, err := repo.CreateVideoProject(userID, &req)
	if err != nil {
		return err.ToHttpRes()
	}

	GenerateVideo(project.ID, *project, repo)

	projectDetail, err := repo.GetVideoProjectDetail(project.ID, userID)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Status = http.StatusOK
	resp.Message = "data created successfully"
	resp.Data = projectDetail

	return
}

func GetVideoProjectList(userID int, repo VideoProjectRepoInterface) (resp response.HttpRes) {
	projects, err := repo.GetVideoProjectList(userID)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Status = http.StatusOK
	resp.Message = "data retrieved successfully"
	resp.Data = projects

	return
}

func GetVideoProjectDetail(userID, projectID int, repo VideoProjectRepoInterface) (resp response.HttpRes) {
	project, err := repo.GetVideoProjectDetail(projectID, userID)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Status = http.StatusOK
	resp.Message = "data retrieved successfully"
	resp.Data = project

	return
}

func GenerateVideo(projectID int, project entity.VideoProject, repo VideoProjectRepoInterface) {
	storyboards, err := gemini_service.GenerateStoryboard(gemini_service.GenerateStoryboardRequest{
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
		// goroutine here
		func(i int, storyboard gemini_service.Storyboard) {
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

			// goroutine here
			//func(sceneID int, storyBoard gemini_service.Storyboard) {
			illustrationUrl, err := dalle_service.GenerateIllustration(storyboard.Illustration, dalle_service.GenerateSize1, 0)
			if err != nil {
				log.Print(err.Message)
				return
			}

			log.Println("generated illustration (sceneID: ", sceneID, ")")

			if err != nil {
				log.Println("error update scene (sceneID: ", sceneID, ")")
				log.Print(err.Message)
				return
			}
			//}(sceneID, storyboard)

			// goroutine here
			//func(sceneID int, storyBoard gemini_service.Storyboard) {
			voiceUrl, err := elevenlabs_service.Generate(storyboard.Narration)
			if err != nil {
				log.Print(err.Message)
				return
			}

			log.Println("generated voice (sceneID: ", sceneID, ")")

			_, err = repo.UpdateScene(sceneID, &entity.Scene{
				IllustrationUrl: illustrationUrl,
				VoiceUrl:        voiceUrl,
			})
			if err != nil {
				log.Println("error update scene (sceneID: ", sceneID, ")")
				log.Print(err.Message)
				return
			}
			//}(sceneID, storyboard)
		}(i, storyboard)
	}
}
