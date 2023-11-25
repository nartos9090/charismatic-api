package pagination

type Pagination struct {
	Page   int `query:"page" json:"page" db:"page"`
	Limit  int `query:"per_page" json:"per_page" db:"limit"`
	Offset int `json:"-" db:"offset"`
	Total  int `json:"total"`
}
