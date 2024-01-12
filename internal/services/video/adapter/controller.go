package adapter

import (
	"go-api-echo/external/services/dalle_service"
	"go-api-echo/external/services/elevenlabs_service"
	"go-api-echo/external/services/gemini_service"
	"go-api-echo/internal/pkg/helpers/response"
	"net/http"
)

func GenerateVideo(req GenerateVideoReq) (resp response.HttpRes) {
	scenes, err := gemini_service.Generate(gemini_service.GenerateRequest(req))
	if err != nil {
		return err.ToHttpRes()
	}

	for _, scene := range scenes {
		illustration, err := dalle_service.Generate(scene.Illustration, dalle_service.GenerateSize1)
		if err != nil {
			continue
		}
		scene.Illustration = illustration

		voice := elevenlabs_service.Generate(scene.Narration)
		scene.Voice = voice
	}

	resp.Status = http.StatusOK
	resp.Message = "data deleted"
	resp.Data = scenes

	return
}
