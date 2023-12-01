package adapter

import (
	"go-api-echo/internal/pkg/crypto"
	errors "go-api-echo/internal/pkg/helpers/errors"
	response "go-api-echo/internal/pkg/helpers/response"
	"go-api-echo/internal/pkg/jwt"
	"net/http"
)

func HandleLogin(req LoginReq, repo AuthRepoInterface) (resp response.HttpRes) {
	admin, err := repo.GetAdminProfile(req.Email)
	if err != nil {
		return err.ToHttpRes()
	}
	if ok := crypto.Match(req.Password, admin.Password, admin.PasswordSalt); !ok {
		err := *errors.UnauthorizedError
		err.AddError("invalid credential")
		return err.ToHttpRes()
	}

	tokenData := jwt.TokenData{
		ID:       admin.ID,
		FullName: admin.FullName,
		Role:     "admin",
	}
	token, err := jwt.GenerateJWT(tokenData)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Data = LoginRes{
		Token:     token,
		TokenData: tokenData,
	}
	resp.Status = http.StatusOK
	resp.Message = "login success"

	return
}
