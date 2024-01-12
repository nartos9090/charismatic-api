package video_infra

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-api-echo/internal/pkg/validator"
	"go-api-echo/internal/services/video/adapter"
)

func VideoRoute(g *echo.Group) {
	r := g.Group("/video")
	defer fmt.Printf(":: Route /employee created\n")

	r.POST("/generate", func(c echo.Context) error {
		var req adapter.GenerateVideoReq
		_ = c.Bind(&req)

		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		resp := adapter.GenerateVideo(req)

		return c.JSON(resp.Status, resp)
	})
}
