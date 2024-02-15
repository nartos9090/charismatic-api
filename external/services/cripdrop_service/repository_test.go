package cripdrop_service

import (
	"encoding/json"
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

func TestReplaceBackground(t *testing.T) {
	config.GlobalEnv.ClipdropConf.ApiKey = ApiKey

	c := CripdropServiceInterface{}

	imagePath := `test.jpg`

	imageResultPath, err := c.ReplaceBackground(imagePath, ``, `mountain view with river and sun`)
	if err != nil {
		println(json.Marshal(err.Message))
		t.Errorf("TestReplaceBackground failed")
		return
	}
	t.Logf(`imagePath: %s`, imageResultPath)
	t.Logf(`TestReplaceBackground success`)
}
