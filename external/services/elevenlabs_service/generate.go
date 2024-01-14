package elevenlabs_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const VOIDE_ID = `ZQe5CZNOzWyzPSCn5a3c`

const MODEL_ID = `eleven_multilingual_v1`

func Generate(text string) (string, *helpers_errors.Error) {
	url := "https://api.elevenlabs.io/v1/text-to-speech/" + VOIDE_ID

	// Define the request payload
	payload := map[string]interface{}{
		`model_id`: MODEL_ID,
		`text`:     text,
		`voice_settings`: map[string]interface{}{
			"similarity_boost":  1,
			"stability":         1,
			"style":             1,
			"use_speaker_boost": true,
		},
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))

	req.Header.Add(`Accept`, `audio/mpeg`)
	req.Header.Add(`xi-api-key`, config.GlobalEnv.ElevenLabsConf.ApiKey)
	req.Header.Add(`Content-Type`, `application/json`)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fileName := filepath.Join(`public/voice`, strconv.Itoa(int(time.Now().Unix()))+".mp3")

	// Create the folder if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating directory`
		return "", commonError
	}

	// Save the file locally
	err = ioutil.WriteFile(fileName, body, 0755)
	if err != nil {
		fmt.Println(err)
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error saving file`
		return "", commonError
	}

	fmt.Println("Voice saved to", fileName)

	return fileName, nil
}
