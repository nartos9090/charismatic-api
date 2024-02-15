package background_removal

type BackgroundRemovalRepoInterface interface {
	RemoveBackground(imagePath string, maskPath string, prompt string) (string, error)
}
