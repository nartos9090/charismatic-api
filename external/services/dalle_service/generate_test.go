package dalle_service

import (
	"go-api-echo/config"
	"testing"
)

const API_KEY = `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`

func TestGenerate(t *testing.T) {
	config.GlobalEnv.DalleConf.ApiKey = API_KEY
	prompt := "mountain view"
	size := GenerateSize2
	t.Logf("prompt: %s", prompt)

	result, err := GenerateIllustration(prompt, size, 0)
	if err != nil {
		t.Errorf("error: %s", err.Message)
		return
	}

	t.Logf("result: %s", result)
}
