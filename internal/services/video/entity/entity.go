package entity

type VideoProject struct {
	ID           int    `json:"id"`
	UserID       int    `json:"user_id"`
	ProductTitle string `json:"product_title"`
	BrandName    string `json:"brand_name"`
	ProductType  string `json:"product_type"`
	MarketTarget string `json:"market_target"`
	Superiority  string `json:"superiority"`
	Duration     int    `json:"duration"`
}

type Scene struct {
	ID              int    `json:"id"`
	VideoProjectID  int    `json:"video_project_id"`
	Sequence        int    `json:"sequence"`
	Title           string `json:"title"`
	Narration       string `json:"narration"`
	Illustration    string `json:"illustration"`
	IllustrationUrl string `json:"illustration_url"`
	VoiceUrl        string `json:"voice_url"`
}
