package background_removal

import "go-api-echo/internal/pkg/helpers/helpers_errors"

type BackgroundRemovalRepoInterface interface {
	RemoveBackground(imagePath string, maskPath string, prompt string) (string, *helpers_errors.Error)
}
