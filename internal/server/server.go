package server

import (
	"go-api-echo/config"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	auth_infra "go-api-echo/internal/services/auth/infra"
)

func InitServer(port string) {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Echo API is running")
	})

	routes := e.Group("/v1")
	// Public routes goes here
	auth_infra.AuthRoute(routes)

	private := routes.Group("")
	private.Use(echojwt.JWT([]byte(config.GlobalEnv.JWTSecret)))
	// Private routes goes here

	e.Logger.Fatal(e.Start(port))
}
