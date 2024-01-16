package adapter

type (
	User struct {
		ID           int    `json:"id" db:"id"`
		Email        string `json:"email" db:"email"`
		FullName     string `json:"fullname" db:"fullname"`
		Password     string `json:"-" db:"passwd"`
		Picture      string `json:"picture" db:"picture"`
		PasswordSalt string `json:"-" db:"passwdSalt"`
		Provider     string `json:"provider" db:"provider"`
		ProviderID   string `json:"provider_id" db:"provider_id"`
	}
)
