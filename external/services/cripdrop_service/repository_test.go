package cripdrop_service

import (
	"go-api-echo/config"
	"testing"
)

const ApiKey = ``

func TestRemoveBackground(t *testing.T) {
	config.GlobalEnv.ClipdropConf.ApiKey = ApiKey

	c := CripdropServiceInterface{}

	imagePath := `test.jpg`

	imageResultPath, err := c.RemoveBackground(imagePath, ``, ``)
	if err != nil {
		println(err.Message)
		t.Errorf("TestRemoveBackground failed")
		return
	}
	t.Logf(`imagePath: %s`, imageResultPath)
	t.Logf(`TestRemoveBackground success`)
}
