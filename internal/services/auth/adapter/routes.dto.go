package adapter

import "go-api-echo/internal/pkg/jwt"

type (
	LoginReq struct {
		Email    string `form:"email" validate:"required"`
		Password string `form:"password" validate:"required"`
	}

	LoginGoogleByAccessTokenReq struct {
		AccessToken string `json:"access_token" validate:"required"`
	}

	LoginGoogleByIdTokenReq struct {
		IdToken string `json:"id_token" validate:"required"`
	}

	LoginRes struct {
		Token   string `json:"token"`
		Picture string `json:"picture"`
		Email   string `json:"email"`
		jwt.TokenData
	}
)
