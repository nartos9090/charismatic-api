package adapter

import (
	"go-api-echo/external/services/googleauth_service"
	"go-api-echo/internal/pkg/crypto"
	errors "go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/pkg/helpers/response"
	"go-api-echo/internal/pkg/jwt"
	"net/http"
)

func HandleLogin(req LoginReq, repo AuthRepoInterface) (resp response.HttpRes) {
	user, err := repo.GetUser(req.Email)
	if err != nil {
		return err.ToHttpRes()
	}

	if user.Password == "" {
		err := *errors.UnauthorizedError
		err.AddError("invalid credential")
		return err.ToHttpRes()
	}

	if ok := crypto.Match(req.Password, user.Password, user.PasswordSalt); !ok {
		err := *errors.UnauthorizedError
		err.AddError("invalid credential")
		return err.ToHttpRes()
	}

	tokenData := jwt.TokenData{
		ID:       user.ID,
		FullName: user.FullName,
		Role:     "user",
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

func HandleGoogleLoginByAccessToken(req LoginGoogleByAccessTokenReq, repo AuthRepoInterface) (resp response.HttpRes) {
	userInfo, err := googleauth_service.ValidateAccessToken(req.AccessToken)
	if err != nil {
		return err.ToHttpRes()
	}

	user, err := repo.GetUser(userInfo.Email)
	if err != nil {
		if err.HttpCode == http.StatusNotFound {
			user = &User{
				Email:      userInfo.Email,
				FullName:   userInfo.Name,
				Picture:    userInfo.Picture,
				Provider:   `google`,
				ProviderID: userInfo.Id,
			}
			user, err = repo.CreateUser(user)
			if err != nil {
				return err.ToHttpRes()
			}
		}
		return err.ToHttpRes()
	}

	tokenData := jwt.TokenData{
		ID:       user.ID,
		FullName: user.FullName,
		Role:     "user",
	}

	token, err := jwt.GenerateJWT(tokenData)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Data = LoginRes{
		Token:     token,
		TokenData: tokenData,
		Email:     userInfo.Email,
		Picture:   user.Picture,
	}
	resp.Status = http.StatusOK
	resp.Message = "login success"

	return
}

func HandleGoogleLoginByIDToken(req LoginGoogleByIdTokenReq, repo AuthRepoInterface) (resp response.HttpRes) {
	payload, err := googleauth_service.ValidateIdToken(req.IdToken)
	if err != nil {
		return err.ToHttpRes()
	}

	email := payload.Claims[`email`].(string)
	name := payload.Claims[`name`].(string)
	picture := payload.Claims[`picture`].(string)
	id := payload.Claims[`sub`].(string)

	user, err := repo.GetUser(email)
	if err != nil {
		if err.HttpCode == http.StatusNotFound {
			user = &User{
				Email:      email,
				FullName:   name,
				Picture:    picture,
				Provider:   `google`,
				ProviderID: id,
			}
			user, err = repo.CreateUser(user)
			if err != nil {
				return err.ToHttpRes()
			}
		}
		return err.ToHttpRes()
	}

	tokenData := jwt.TokenData{
		ID:       user.ID,
		FullName: user.FullName,
		Role:     "user",
	}

	token, err := jwt.GenerateJWT(tokenData)
	if err != nil {
		return err.ToHttpRes()
	}

	resp.Data = LoginRes{
		Token:     token,
		TokenData: tokenData,
		Email:     email,
		Picture:   user.Picture,
	}
	resp.Status = http.StatusOK
	resp.Message = "login success"

	return
}
