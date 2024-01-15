package entity

type History struct {
	CreatedAt string `json:"created_at" db:"created_at"`
	Title     string `json:"title" db:"title"`
	Type      string `json:"type" db:"type"`
	ID        int    `json:"id" db:"id"`
}
