package infra

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-api-echo/internal/pkg/db/mysql"
	"go-api-echo/internal/pkg/validator"
	adapter "go-api-echo/internal/services/auth/adapter"
	"time"
)

var CONTEXT_TIMEOUT = 15 * time.Second

func AuthRoute(g *echo.Group) {
	r := g.Group("/auth")
	defer fmt.Printf(":: Route /auth created\n")

	r.POST("/login", func(c echo.Context) error {
		var req adapter.LoginReq
		_ = c.Bind(&req)

		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		timeout := CONTEXT_TIMEOUT
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := &AuthRepo{
			ctx: ctx,
			db:  db_mysql.Db,
		}
		resp := adapter.HandleLogin(req, repo)

		return c.JSON(resp.Status, resp)
	})

	r.POST("/login/google", func(c echo.Context) error {
		var req adapter.LoginGoogleByAccessTokenReq
		_ = c.Bind(&req)

		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		timeout := CONTEXT_TIMEOUT
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := &AuthRepo{
			ctx: ctx,
			db:  db_mysql.Db,
		}
		resp := adapter.HandleGoogleLoginByAccessToken(req, repo)

		return c.JSON(resp.Status, resp)
	})

	r.POST("/login/google/id-token", func(c echo.Context) error {
		var req adapter.LoginGoogleByIdTokenReq
		_ = c.Bind(&req)

		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		timeout := CONTEXT_TIMEOUT
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := &AuthRepo{
			ctx: ctx,
			db:  db_mysql.Db,
		}
		resp := adapter.HandleGoogleLoginByIDToken(req, repo)

		return c.JSON(resp.Status, resp)
	})
}
