package dalle_service

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var GenerateSize1 = "256x256"
var GenerateSize2 = "512x512"
var GenerateSize3 = "1024x1024"

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

func readImageFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return content, nil
}
