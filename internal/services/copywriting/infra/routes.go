package infra

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-api-echo/internal/pkg/db/mysql"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/pkg/jwt"
	"go-api-echo/internal/pkg/validator"
	"go-api-echo/internal/services/copywriting/adapter"
	"strconv"
)

func CopywritingRoute(g *echo.Group) {
	r := g.Group("/copywriting-project")
	defer fmt.Printf(":: route /copywriting-project created\n")

	r.POST("/create-sync", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		var req adapter.CreateCopywritingReq
		_ = c.Bind(&req)

		file, err := c.FormFile("product_image")
		if err != nil {
			comErr := helpers_errors.BadRequestError
			comErr.AddError("invalid product_image")
			resp := comErr.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		req.ProductImage, err = file.Open()
		if err != nil {
			comErr := helpers_errors.BadRequestError
			comErr.AddError("invalid product_image")
			resp := comErr.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		repo := CopywritingProjectRepository{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.CreateCopywritingProjectSync(user.ID, req, repo)

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

		repo := CopywritingProjectRepository{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.GetCopywritingDetail(user.ID, id, repo)

		return c.JSON(resp.Status, resp)
	})

	r.GET("/list", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		repo := CopywritingProjectRepository{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		//resp := adapter.GetCopywritingList(user.ID, repo)
		resp := adapter.GetCopywritingList(user.ID, repo)

		return c.JSON(resp.Status, resp)
	})
}
