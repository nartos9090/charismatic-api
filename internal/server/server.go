package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
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
	video_infra.VideoRoute(routes)

	routes.Static("/public", "public")

	e.Logger.Fatal(e.Start(port))
}
