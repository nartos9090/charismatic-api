package jwt

import (
	"go-api-echo/config"
	errors "go-api-echo/internal/pkg/helpers/helpers_errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
)

type TokenData struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Role     string `json:"role"`
}

type JwtClaims struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateJWT(data TokenData) (string, *errors.Error) {
	claims := &JwtClaims{
		ID:       data.ID,
		FullName: data.FullName,
		Role:     data.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GlobalEnv.JWTSecret))
	if err != nil {
		log.Print(err.Error())
		return "", errors.InternalServerError
	}

	return t, nil
}

func Authorize(c echo.Context, roles ...string) *JwtClaims {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JwtClaims)

	if claims.ID == 0 {
		resp := errors.UnauthorizedError.ToHttpRes()
		c.JSON(resp.Status, resp)
		return nil
	}
	if len(roles) > 0 {
		authenticated := false

		for _, role := range roles {
			if role == claims.Role {
				authenticated = true
				break
			}
		}

		if !authenticated {
			resp := errors.ForbiddenError.ToHttpRes()
			c.JSON(resp.Status, resp)
			return nil
		}
	}

	return claims
}
