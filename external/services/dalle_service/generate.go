package dalle_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const FOLDER_PATH = `public/images`

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
func Generate(prompt string, size string, iterate int) (string, *helpers_errors.Error) {
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
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}

	// Make the API call
	request, err := http.NewRequest(`POST`, apiEndpoint, bytes.NewBuffer(payloadBytes))
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating external client`
		return "", commonError
	}

	request.Header.Set(`Content-Type`, `application/json`)
	request.Header.Set(`Authorization`, `Bearer `+config.GlobalEnv.DalleConf.ApiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		commonError := helpers_errors.BadGatewayError
		commonError.Message = `error making external request`
		return "", commonError
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusUnauthorized {
		externalError := helpers_errors.BadGatewayError
		externalError.Message = `error on external request (unauthorized)`
		return "", externalError
	}

	if response.StatusCode == http.StatusTooManyRequests {
		if iterate > 10 {
			externalError := helpers_errors.BadGatewayError
			externalError.Message = `error on external request (too many requests)`
			return "", externalError
		} else {
			iterate++
			return Generate(prompt, size, iterate)
		}
	}

	if response.StatusCode != http.StatusOK {
		externalError := helpers_errors.BadGatewayError
		externalError.Message = `error on external request`
		return "", externalError
	}

	// Decode the response JSON
	var dalleResponse DALLEResponse
	if err := json.NewDecoder(response.Body).Decode(&dalleResponse); err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `internal server error`
		return "", commonError
	}

	// Extract the image URL from the response
	imageURL := dalleResponse.Data[0].URL

	// Download and save the image locally
	fileName := filepath.Join(FOLDER_PATH, strconv.Itoa(int(dalleResponse.Created))+`.png`)
	if err := downloadImage(imageURL, fileName); err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `internal server error`
		return "", commonError
	}

	fmt.Println("Illustration saved to", fileName)

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
