package elevenlabs_service

import (
	"go-api-echo/config"
	"testing"
)

const API_KEY = ``

func TestGenerate(t *testing.T) {
	config.GlobalEnv.ElevenLabsConf.ApiKey = API_KEY
	text := `AquaVita merupakan pilihan tepat untuk gaya hidup aktif dan sibuk, dengan desain kemasan yang elegan dan praktis.`
	t.Logf("prompt: %s", text)

	result, err := Generate(text)
	if err != nil {
		t.Errorf("error: %s", err.Message)
		return
	}

	t.Logf("result: %s", result)
}
