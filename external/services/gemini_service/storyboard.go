package gemini_service

import (
	"context"
	"github.com/google/generative-ai-go/genai"
	"go-api-echo/config"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"google.golang.org/api/option"
	"log"
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
		"Buat storyboard untuk membuat video iklan promosi produk dengan detail produk di atas. Buat narasi iklan yang menarik, dengan masing-masing perkiraan durasi 10 detik. Tambahkan judul ilustrasi gambar yang cocok. Tuliskan masing-masing jawaban dalam satu baris dengan format output berikut\nJudul Adegan:\nTeks Narasi Iklan:\nTeks Ilustrasi Gambar:"
}

func GenerateStoryboard(req GenerateStoryboardRequest) ([]Storyboard, *helpers_errors.Error) {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(config.GlobalEnv.GeminiConf.ApiKey))

	if err != nil {
		commonErr := *helpers_errors.InternalServerError
		commonErr.AddError("internal server error")
		return nil, &commonErr
	}

	defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	resp, err := model.GenerateContent(ctx, genai.Text(generateStoryboardPrompt(req)))
	if err != nil {
		log.Fatal(err)
	}

	text := parseResponse(resp)
	text = strings.Replace(text, "**", "", -1)

	print(text)

	// TODO: fix error
	return parseStoryboardScenes(text), nil
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
