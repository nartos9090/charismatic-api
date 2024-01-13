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

	result, err := Generate(prompt, size)
	if err != nil {
		t.Errorf("error: %s", err.Message)
		return
	}

	t.Logf("result: %s", result)
}
