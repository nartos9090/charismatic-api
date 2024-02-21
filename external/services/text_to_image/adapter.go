package text_to_image

import "go-api-echo/internal/pkg/helpers/helpers_errors"

type TextToImageRepoInterface interface {
	Generate(prompt string) (string, *helpers_errors.Error)
}
