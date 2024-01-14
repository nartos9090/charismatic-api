package server

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"go-api-echo/config"
	errors "go-api-echo/internal/pkg/helpers/helpers_errors"
	customjwt "go-api-echo/internal/pkg/jwt"
	auth_infra "go-api-echo/internal/services/auth/infra"
	copywriting_infra "go-api-echo/internal/services/copywriting/infra"
	video_infra "go-api-echo/internal/services/video/infra"
	"net/http"
)

func InitServer(port string) {
	fmt.Printf(":: Perparing routes\n\n")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo API is running")
	})

	routes := e.Group("/v1")
	auth_infra.AuthRoute(routes)

	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(customjwt.JwtClaims)
		},
		SigningKey: []byte(config.GlobalEnv.JWTSecret),
		ErrorHandler: func(c echo.Context, err error) error {
			res := errors.UnauthorizedError
			res.Errors = append(res.Errors, err.Error())
			return c.JSON(res.HttpCode, res.ToHttpRes())
		},
	}

	private := routes.Group("")
	private.Use(echojwt.WithConfig(jwtConfig))
	video_infra.VideoRoute(private)
	copywriting_infra.CopywritingRoute(private)

	routes.Static("/public", "public")

	e.Logger.Fatal(e.Start(port))
}
