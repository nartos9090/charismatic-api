package infra

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-api-echo/internal/pkg/db/mysql"
	"go-api-echo/internal/pkg/jwt"
	"go-api-echo/internal/services/history/adapter"
)

func HistoryRoute(g *echo.Group) {
	r := g.Group("/history")
	defer fmt.Printf(":: route /history created\n")

	r.GET("", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		repo := HistoryRepoInterface{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.GetHistory(user.ID, repo)

		return c.JSON(resp.Status, resp)
	})
}
