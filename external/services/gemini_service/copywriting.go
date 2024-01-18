package gemini_service

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"google.golang.org/api/option"
	"io"
	"log"
	"mime/multipart"
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
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GlobalEnv.GeminiConf.ApiKey))
	if err != nil {
		log.Fatal(err)
		comErr := *helpers_errors.InternalServerError
		comErr.AddError("internal server error")
		return "", &comErr
	}
	defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	// convert product image into byte
	var fileBytes []byte

	// Read the file content into a byte slice
	buffer := make([]byte, 1024)
	for {
		n, err := req.ProductImage.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			comErr := *helpers_errors.InternalServerError
			comErr.AddError("internal server error")
			return "", &comErr
		}
		fileBytes = append(fileBytes, buffer[:n]...)
	}

	prompt := []genai.Part{
		genai.ImageData("jpeg", fileBytes),
		genai.Text(generateCopywritingPrompt(req)),
	}
	resp, err := model.GenerateContent(ctx, prompt...)

	if err != nil {
		log.Fatal(err)
		comErr := *helpers_errors.BadGatewayError
		comErr.AddError("bad gateway")
		return "", &comErr
	}

	return parseResponse(resp), nil
}
