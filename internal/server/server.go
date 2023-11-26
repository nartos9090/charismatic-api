package server

import (
	"fmt"
	"go-api-echo/config"
	"net/http"

	"go-api-echo/internal/pkg/helpers/errors"
	customjwt "go-api-echo/internal/pkg/jwt"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	auth_infra "go-api-echo/internal/services/auth/infra"
	employee_infra "go-api-echo/internal/services/employee/infra"
)

func InitServer(port string) {
	fmt.Printf(":: Perparing routes\n\n")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo API is running")
	})

	routes := e.Group("/v1")
	// Public routes goes here
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
	// Private routes goes here
	employee_infra.EmployeeRoute(private)

	e.Logger.Fatal(e.Start(port))
}
