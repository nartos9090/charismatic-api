package adapter

import (
	"go-api-echo/external/services/dalle_service"
	"go-api-echo/external/services/elevenlabs_service"
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

	wg.Add(len(scenes) * 2)

	for i, _scene := range scenes {
		go func(i int, scene gemini_service.Scene) {
			defer wg.Done()
			url, err := dalle_service.Generate(scene.Illustration, dalle_service.GenerateSize1)
			if err != nil {
				resp = err.ToHttpRes()
				return
			}
			scenes[i].IllustrationUrl = url
		}(i, _scene)

		go func(i int, scene gemini_service.Scene) {
			defer wg.Done()
			voice, err := elevenlabs_service.Generate(scene.Narration)
			if err != nil {
				resp = err.ToHttpRes()
				return
			}
			scenes[i].VoiceUrl = voice
		}(i, _scene)
	}

	wg.Wait()

	resp.Status = http.StatusOK
	resp.Message = "data created successfully"
	resp.Data = scenes

	return
}
