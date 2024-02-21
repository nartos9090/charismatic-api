package gemini_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"io/ioutil"
	"net/http"
	"strings"
)

type GenerateStoryboardRequest struct {
	ProductTitle string `json:"product_title"`
	BrandName    string `json:"brand_name"`
	ProductType  string `json:"product_type"`
	MarketTarget string `json:"market_target"`
	Superiority  string `json:"superiority"`
	Duration     int    `json:"duration"`
}

type Storyboard struct {
	Order           int    `json:"order"`
	Title           string `json:"title"`
	Narration       string `json:"narration"`
	Illustration    string `json:"illustration"`
	IllustrationUrl string `json:"illustration_url"`
	Voice           string `json:"voice"`
	VoiceUrl        string `json:"voice_url"`
}

func generateStoryboardPrompt(req GenerateStoryboardRequest) string {
	return "Judul Produk: " + req.ProductTitle + "\n\n" +
		"Nama Brand: " + req.BrandName + "\n\n" +
		"Jenis Produk: " + req.ProductType + "\n\n" +
		"Target Pasar: " + req.MarketTarget + "\n\n" +
		"Keunggulan: " + req.Superiority + "\n\n" +
		// TODO: fix duration prompt
		//"Duration: " + strconv.Itoa(int(req.Duration/10)) + "\n\n" +
		"Buat 5 storyboard untuk membuat video iklan promosi produk dengan detail produk di atas. Buat narasi iklan yang menarik, dengan masing-masing perkiraan durasi 10 detik. Tambahkan judul ilustrasi gambar yang cocok. Tuliskan masing-masing jawaban dalam satu baris dengan format output berikut\nJudul Adegan:\nTeks Narasi Iklan:\nTeks Ilustrasi Gambar:"
}

func GenerateStoryboard(req GenerateStoryboardRequest) ([]Storyboard, *helpers_errors.Error) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.0-pro:generateContent?key=" + config.GlobalEnv.GeminiConf.ApiKey

	payload := []byte(`{
		"contents": [
			{
				"role": "user",
				"parts": [
					{
						"text": "` + generateStoryboardPrompt(req) + `"
					}
				]
			}
		],
		"generationConfig": {
			"temperature": 0.9,
			"topK": 1,
			"topP": 1,
			"maxOutputTokens": 2048,
			"stopSequences": []
		},
		"safetySettings": [
			{
				"category": "HARM_CATEGORY_HARASSMENT",
				"threshold": "BLOCK_MEDIUM_AND_ABOVE"
			},
			{
				"category": "HARM_CATEGORY_HATE_SPEECH",
				"threshold": "BLOCK_MEDIUM_AND_ABOVE"
			},
			{
				"category": "HARM_CATEGORY_SEXUALLY_EXPLICIT",
				"threshold": "BLOCK_MEDIUM_AND_ABOVE"
			},
			{
				"category": "HARM_CATEGORY_DANGEROUS_CONTENT",
				"threshold": "BLOCK_MEDIUM_AND_ABOVE"
			}
		]
	}`)

	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		commonErr := *helpers_errors.InternalServerError
		commonErr.AddError("error on external request")
		return nil, &commonErr
	}

	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		commonErr := *helpers_errors.InternalServerError
		commonErr.AddError("error on external request")
		return nil, &commonErr
	}
	defer resp.Body.Close()

	fmt.Println("HTTP Status Code:", resp.StatusCode)

	// Read the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error reading external response`
		return nil, commonError
	}

	var response Response
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error unmarshalling external response`
		return nil, commonError
	}

	// TODO: fix error
	return parseStoryboardScenes(response.Candidates[0].Content.Parts[0].Text), nil
}

func parseStoryboardScenes(result string) []Storyboard {
	lines := strings.Split(result, "\n")

	var scenes []Storyboard

	sceneNumber := 0

	for _, line := range lines {
		if strings.HasPrefix(line, "Judul Adegan: ") {
			sceneNumber++
			// TODO: trim
			scene := Storyboard{
				Order: sceneNumber,
				Title: strings.Replace(line, "Judul Adegan: ", "", 1),
			}
			scenes = append(scenes, scene)
		} else if strings.HasPrefix(line, "Teks Narasi Iklan: ") {
			// TODO: trim
			scenes[len(scenes)-1].Narration = strings.Replace(line, "Teks Narasi Iklan: ", "", 1)
		} else if strings.HasPrefix(line, "Teks Ilustrasi Gambar: ") {
			// TODO: trim
			scenes[len(scenes)-1].Illustration = strings.Replace(line, "Teks Ilustrasi Gambar: ", "", 1)
		}
	}

	return scenes
}
