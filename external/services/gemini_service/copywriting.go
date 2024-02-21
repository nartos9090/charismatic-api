package gemini_service

import (
	"bytes"
	"encoding/json"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

type GenerateCopywritingRequest struct {
	BrandName    string
	MarketTarget string
	ProductImage multipart.File
	Superiority  string
	Duration     int
}

func generateCopywritingPrompt(req GenerateCopywritingRequest) string {
	return "Buatkan teks copywriting mengenai gambar di atas yang bernama " + req.BrandName +
		" dengan target pasar " + req.MarketTarget +
		" dengan maksimal 1000 huruf dan keunggulannya sebagai berikut tanpa emoji. " + req.Superiority
}

func GenerateCopywriting(req GenerateCopywritingRequest) (string, *helpers_errors.Error) {
	url := "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.0-pro:generateContent?key=" + config.GlobalEnv.GeminiConf.ApiKey

	payload := []byte(`{
		"contents": [
			{
				"role": "user",
				"parts": [
					{
						"text": "` + generateCopywritingPrompt(req) + `"
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
		return ``, &commonErr
	}

	httpReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(httpReq)
	if err != nil {
		commonErr := *helpers_errors.InternalServerError
		commonErr.AddError("error on external request")
		return ``, &commonErr
	}
	defer resp.Body.Close()

	// Read the response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error reading external response`
		return "", commonError
	}

	var response Response
	err = json.Unmarshal([]byte(responseBody), &response)
	if err != nil {
		commonError := helpers_errors.InternalServerError
		commonError.Message = `error unmarshalling external response`
		return "", commonError
	}

	return response.Candidates[0].Content.Parts[0].Text, nil
}
