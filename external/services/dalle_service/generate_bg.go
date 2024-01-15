package dalle_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

// GenerateBackground function makes an API call to DALL-E 2, downloads the image, and saves it locally.
func GenerateBackground(imagePath string, maskPath string, prompt string) (string, *helpers_errors.Error) {
	apiEndpoint := "https://api.openai.com/v1/images/edits"

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add image file to the request
	imageFile, err := os.Open(imagePath)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}
	defer imageFile.Close()

	imagePart, err := writer.CreateFormFile("image", "image.jpg")
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}
	_, err = io.Copy(imagePart, imageFile)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}

	// Add mask file to the request
	maskFile, err := os.Open(maskPath)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}
	defer maskFile.Close()

	maskPart, err := writer.CreateFormFile("mask", "mask.jpg")
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}
	_, err = io.Copy(maskPart, maskFile)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}

	// Add prompt to the request
	writer.WriteField("prompt", prompt)
	writer.WriteField("n", "1")
	writer.WriteField("model", "dall-e-2")
	writer.WriteField("size", GenerateSize3)

	// Close the multipart writer to finalize the request
	err = writer.Close()
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}

	// Make the API call
	request, err := http.NewRequest(`POST`, apiEndpoint, body)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating external client`
		return "", commonError
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
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
		externalError := helpers_errors.BadGatewayError
		externalError.Message = `error on external request (too many requests)`
		return "", externalError
	}

	fmt.Println(response.StatusCode)
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

	fmt.Println("Image saved to", fileName)

	return fileName, nil
}
