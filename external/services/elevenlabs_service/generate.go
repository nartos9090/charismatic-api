package elevenlabs_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Generate(text string) string {
	url := "https://api.elevenlabs.io/v1/text-to-speech/ZQe5CZNOzWyzPSCn5a3c"

	// Define the request payload
	payload := map[string]interface{}{
		`model_id`: `eleven_multilingual_v1`,
		`text`:     text,
		`voice_settings`: map[string]interface{}{
			"similarity_boost":  123,
			"stability":         123,
			"style":             123,
			"use_speaker_boost": true,
		},
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	result, err := json.Marshal(body)
	if err != nil {
		return ``
	}
	return string(result)
}
