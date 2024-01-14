package infra

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-api-echo/internal/pkg/db/mysql"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/pkg/jwt"
	"go-api-echo/internal/pkg/validator"
	"go-api-echo/internal/services/video/adapter"
	"strconv"
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
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.AddVideoProject(user.ID, req, repo)

		return c.JSON(resp.Status, resp)
	})

	r.POST("/generate-sync", func(c echo.Context) error {
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
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.AddVideoProjectSyncWithDetail(user.ID, req, repo)

		return c.JSON(resp.Status, resp)
	})

	r.GET("/list", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		repo := VideoProjectRepo{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.GetVideoProjectList(user.ID, repo)

		return c.JSON(resp.Status, resp)
	})

	r.GET("/detail/:id", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			comErr := helpers_errors.BadRequestError
			comErr.AddError("invalid id")
			resp := comErr.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		repo := VideoProjectRepo{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.GetVideoProjectDetail(user.ID, id, repo)

		return c.JSON(resp.Status, resp)
	})
}
