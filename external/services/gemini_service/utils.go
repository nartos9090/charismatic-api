package gemini_service

import (
	"fmt"
	"github.com/google/generative-ai-go/genai"
)

func parseResponse(resp *genai.GenerateContentResponse) (text string) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				text += fmt.Sprintf("%s", part)
			}
		}
	}

	return
}

type Response struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
			Role string `json:"role"`
		} `json:"content"`
	} `json:"candidates"`
}
