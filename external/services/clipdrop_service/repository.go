package clipdrop_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

type CripdropServiceInterface struct {
}

const ApiRemoveBackgroundEndpoint = `https://clipdrop-api.co/remove-background/v1`
const ApiReplaceBackgroundEndpoint = `https://clipdrop-api.co/replace-background/v1`
const FolderPath = `public/images`

func (c CripdropServiceInterface) RemoveBackground(imagePath string, maskPath string, prompt string) (string, *helpers_errors.Error) {
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

	imagePart, err := writer.CreateFormFile("image_file", imagePath)
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

	// Close the multipart writer to finalize the request
	err = writer.Close()
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}

	// Make the API call
	request, err := http.NewRequest(`POST`, ApiRemoveBackgroundEndpoint, body)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating external client`
		return "", commonError
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set(`x-api-key`, config.GlobalEnv.ClipdropConf.ApiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		commonError := helpers_errors.BadGatewayError
		commonError.Message = `error making external request`
		return "", commonError
	}

	defer response.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error reading external response`
		return "", commonError
	}

	if response.StatusCode != http.StatusOK {
		fmt.Println(string(responseBody))
		externalError := helpers_errors.BadGatewayError
		externalError.Message = `error on external request`
		return "", externalError
	}

	return saveFileToFolder(responseBody)
}

func (c CripdropServiceInterface) ReplaceBackground(imagePath string, maskPath string, prompt string) (string, *helpers_errors.Error) {
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

	imagePart, err := writer.CreateFormFile("image_file", "image.jpg")
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

	writer.WriteField(`prompt`, prompt)

	// Close the multipart writer to finalize the request
	err = writer.Close()
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}

	// Make the API call
	request, err := http.NewRequest(`POST`, ApiReplaceBackgroundEndpoint, body)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating external client`
		return "", commonError
	}

	request.Header.Set("Content-Type", writer.FormDataContentType())
	request.Header.Set(`x-api-key`, config.GlobalEnv.ClipdropConf.ApiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		commonError := helpers_errors.BadGatewayError
		commonError.Message = `error making external request`
		return "", commonError
	}

	defer response.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error reading external response`
		return "", commonError
	}

	if response.StatusCode != http.StatusOK {
		externalError := helpers_errors.BadGatewayError
		externalError.Message = `error on external request`

		type ErrorResponse struct {
			Error string `json:"error"`
		}
		var errorResponse ErrorResponse
		err = json.Unmarshal(responseBody, &errorResponse)
		if err != nil {
			externalError.AddError()
			return "", externalError
		}
		externalError.AddError(errorResponse.Error)
		return "", externalError
	}

	return saveFileToFolder(responseBody)
}
