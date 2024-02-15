package cripdrop_service

import (
	"fmt"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func saveFileToFolder(responseBody []byte) (string, *helpers_errors.Error) {
	if err := os.MkdirAll(FolderPath, os.ModePerm); err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error creating folder`
		return "", commonError
	}

	// get current timestamp
	fileName := filepath.Join(FolderPath, strconv.Itoa(int(time.Now().Unix()))+`.png`)
	// write the response to a file
	err := ioutil.WriteFile(fileName, responseBody, 0644)
	if err != nil {
		fmt.Println(err)
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error writing response`
		return "", commonError
	}

	return fileName, nil
}
