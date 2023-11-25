package auth_adapter

import "go-api-echo/internal/pkg/jwt"

type (
	LoginReq struct {
		Email    string `form:"email" validate:"required"`
		Password string `form:"password" validate:"required"`
	}

	LoginRes struct {
		Token string `json:"token"`
		jwt.TokenData
	}
)
