package infra

import (
	"fmt"
	"github.com/labstack/echo/v4"
	db_mysql "go-api-echo/internal/pkg/db/mysql"
	"go-api-echo/internal/pkg/jwt"
	"go-api-echo/internal/pkg/validator"
	"go-api-echo/internal/services/video/adapter"
)

func VideoRoute(g *echo.Group) {
	r := g.Group("/video-project")
	defer fmt.Printf(":: route /video-project created\n")

	r.POST("/generate", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		var req adapter.GenerateVideoReq
		_ = c.Bind(&req)

		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		repo := VideoProjectRepo{
			ctx: c.Request().Context(),
			db:  db_mysql.Db,
		}

		resp := adapter.AddVideoProject(user.ID, req, repo)

		return c.JSON(resp.Status, resp)
	})
}
