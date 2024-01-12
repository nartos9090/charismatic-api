package adapter

import (
	"fmt"
	"go-api-echo/external/services/dalle_service"
	"go-api-echo/external/services/gemini_service"
	"go-api-echo/internal/pkg/helpers/response"
	"net/http"
	"sync"
)

func GenerateVideo(req GenerateVideoReq) (resp response.HttpRes) {
	scenes, err := gemini_service.Generate(gemini_service.GenerateRequest(req))
	if err != nil {
		return err.ToHttpRes()
	}

	var wg sync.WaitGroup

	wg.Add(len(scenes))

	for i := range scenes {
		go func(i int) {
			defer wg.Done()
			fmt.Println(scenes[i].Illustration)
			illustration, err := dalle_service.Generate(scenes[i].Illustration, dalle_service.GenerateSize1)
			if err != nil {
				return
			}
			scenes[i].Illustration = illustration

			//voice := elevenlabs_service.Generate(scene.Narration)
			//scene.Voice = voice
		}(i)
	}

	wg.Wait()

	resp.Status = http.StatusOK
	resp.Message = "data created successfully"
	resp.Data = scenes

	return
}
