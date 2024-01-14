package entity

type VideoProject struct {
	ID           int    `json:"id" db:"id"`
	UserID       int    `json:"user_id" db:"user_id"`
	ProductTitle string `json:"product_title" db:"product_title"`
	BrandName    string `json:"brand_name" db:"brand_name"`
	ProductType  string `json:"product_type" db:"product_type"`
	MarketTarget string `json:"market_target" db:"market_target"`
	Superiority  string `json:"superiority" db:"superiority"`
	Duration     int    `json:"duration" db:"duration"`
}

type Scene struct {
	ID              int    `json:"id" db:"id"`
	VideoProjectID  int    `json:"video_project_id" db:"video_project_id"`
	Sequence        int    `json:"sequence" db:"sequence"`
	Title           string `json:"title" db:"title"`
	Narration       string `json:"narration" db:"narration"`
	Illustration    string `json:"illustration" db:"illustration"`
	IllustrationUrl string `json:"illustration_url" db:"illustration_url"`
	VoiceUrl        string `json:"voice_url" db:"voice_url"`
}
