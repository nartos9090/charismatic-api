package jwt

import (
	"go-api-echo/config"
	errors "go-api-echo/internal/pkg/helpers/errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4/middleware"
)

type TokenData struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Picture *string `json:"picture"`
	Role    string  `json:"role"`
}

type JwtClaims struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Picture *string `json:"picture"`
	Role    string  `json:"role"`
	jwt.StandardClaims
}

func GenerateJWT(data TokenData) (string, *errors.Error) {
	claims := &JwtClaims{
		ID:      data.ID,
		Name:    data.Name,
		Picture: data.Picture,
		Role:    data.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(config.GlobalEnv.JWTSecret))

	if err != nil {
		log.Print(err.Error())
		return ``, errors.InternalServerError
	}

	return t, nil
}

func Authorize(c echo.Context, roles ...string) *JwtClaims {
	user := c.Get(`user`).(*jwt.Token)
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
