package auth_adapter

type (
	Admin struct {
		ID           int    `json:"id" db:"id"`
		Email        string `json:"email" db:"email"`
		FullName     string `json:"fullname" db:"fullname"`
		Password     string `json:"-" db:"passwd"`
		PasswordSalt string `json:"-" db:"passwdSalt"`
	}
)
