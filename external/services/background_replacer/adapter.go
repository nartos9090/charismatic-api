package background_replacer

import "go-api-echo/internal/pkg/helpers/helpers_errors"

type BackgroundReplacerRepoInterface interface {
	ReplaceBackground(imagePath string, maskPath string, prompt string) (string, *helpers_errors.Error)
}
