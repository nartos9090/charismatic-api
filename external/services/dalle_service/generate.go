package dalle_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const FOLDER_PATH = `public/images`

// var apiKey = os.Getenv(`DALL_E_API_KEY`)
var apiKey = `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`

var GenerateSize1 = "256x256"
var GenerateSize2 = "512x512"
var GenerateSize3 = "1024x1024"

type DALLEResponse struct {
	Created int64 `json:"created"`
	Data    []struct {
		URL string `json:"url"`
	} `json:"data"`
}

// Generate function makes an API call to DALL-E 2, downloads the image, and saves it locally.
func Generate(prompt string, size string) (string, error) {
	apiEndpoint := "https://api.openai.com/v1/images/generations"

	// Define the request payload
	payload := map[string]interface{}{
		`model`:  `dall-e-2`,
		`prompt`: prompt,
		`n`:      1,
		`size`:   size,
	}

	// Convert payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	// Make the API call
	request, err := http.NewRequest(`POST`, apiEndpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return "", err
	}

	request.Header.Set(`Content-Type`, `application/json`)
	request.Header.Set(`Authorization`, `Bearer `+apiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("DALL-E API returned non-OK status: %d", response.StatusCode)
	}

	// Decode the response JSON
	var dalleResponse DALLEResponse
	if err := json.NewDecoder(response.Body).Decode(&dalleResponse); err != nil {
		return "", err
	}

	// Extract the image URL from the response
	imageURL := dalleResponse.Data[0].URL

	// Download and save the image locally
	fileName := filepath.Join(FOLDER_PATH, strconv.Itoa(int(dalleResponse.Created))+`.png`)
	if err := downloadImage(imageURL, fileName); err != nil {
		return "", err
	}

	return fileName, nil
}

// downloadImage function downloads the image from the given URL and saves it locally.
func downloadImage(url, fileName string) error {
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Create the folder if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(fileName), 0755); err != nil {
		return err
	}

	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}
