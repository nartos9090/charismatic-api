package cripdrop_service

import (
	"bytes"
	"fmt"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type CripdropServiceInterface struct {
}

const ApiEndpoint = `https://clipdrop-api.co/remove-background/v1`
const FolderPath = `public/images`

func (c *CripdropServiceInterface) RemoveBackground(imagePath string, maskPath string, prompt string) (string, *helpers_errors.Error) {
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

	// Close the multipart writer to finalize the request
	err = writer.Close()
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating payload`
		return "", commonError
	}

	// Make the API call
	request, err := http.NewRequest(`POST`, ApiEndpoint, body)
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
		return "", externalError
	}

	// Create the destination imageFile
	uploadPath := "./public/images/" // Change this path as needed
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating folder`
		return "", commonError
	}

	// get current timestamp
	fileName := filepath.Join(FolderPath, strconv.Itoa(int(time.Now().Unix()))+`.png`)
	// write the response to a file
	err = ioutil.WriteFile(fileName, responseBody, 0644)
	if err != nil {
		fmt.Println(err)
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error writing response`
		return "", commonError
	}

	return fileName, nil
}
