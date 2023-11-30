package infra

import (
	"context"
	"fmt"
	"go-api-echo/internal/pkg/db/sqlite"
	"go-api-echo/internal/pkg/jwt"
	"go-api-echo/internal/pkg/validator"
	"go-api-echo/internal/services/employee/adapter"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func EmployeeRoute(g *echo.Group) {
	r := g.Group("/employee")
	defer fmt.Printf(":: Route /employee created\n")

	r.GET("", func(c echo.Context) error {
		admin := jwt.Authorize(c)
		if admin == nil {
			return nil
		}

		var req adapter.ListEmployeeReq
		_ = c.Bind(&req)
		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}
		req.Pagination.Parse()

		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := EmployeeRepo{
			ctx: ctx,
			db:  sqlite.Db,
		}

		resp := adapter.HandleListEmployee(req, &repo)
		return c.JSON(resp.Status, resp)
	})

	r.GET("/:id", func(c echo.Context) error {
		admin := jwt.Authorize(c)
		if admin == nil {
			return nil
		}

		rawId := c.Param("id")
		id, _ := strconv.Atoi(rawId)

		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := EmployeeRepo{
			ctx: ctx,
			db:  sqlite.Db,
		}

		resp := adapter.HandleDetailEmployee(id, &repo)
		return c.JSON(resp.Status, resp)
	})

	r.POST("", func(c echo.Context) error {
		admin := jwt.Authorize(c)
		if admin == nil {
			return nil
		}

		var req adapter.UpsertEmployeeReq
		_ = c.Bind(&req)
		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := EmployeeRepo{
			ctx: ctx,
			db:  sqlite.Db,
		}

		resp := adapter.HandleInsertEmployee(req, &repo)
		return c.JSON(resp.Status, resp)
	})

	r.PUT("/:id", func(c echo.Context) error {
		admin := jwt.Authorize(c)
		if admin == nil {
			return nil
		}

		var req adapter.UpsertEmployeeReq
		_ = c.Bind(&req)
		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := EmployeeRepo{
			ctx: ctx,
			db:  sqlite.Db,
		}

		resp := adapter.HandleUpdateEmployee(req, &repo)
		return c.JSON(resp.Status, resp)
	})

	r.DELETE("/:id", func(c echo.Context) error {
		admin := jwt.Authorize(c)
		if admin == nil {
			return nil
		}

		rawId := c.Param("id")
		id, _ := strconv.Atoi(rawId)

		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := EmployeeRepo{
			ctx: ctx,
			db:  sqlite.Db,
		}

		resp := adapter.HandleDeleteEmployee(id, &repo)
		return c.JSON(resp.Status, resp)
	})

	r.GET("/:id/leave-submission", func(c echo.Context) error {
		admin := jwt.Authorize(c)
		if admin == nil {
			return nil
		}

		rawId := c.Param("id")
		id, _ := strconv.Atoi(rawId)

		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := EmployeeRepo{
			ctx: ctx,
			db:  sqlite.Db,
		}

		resp := adapter.HandleListEmployeeLeave(id, &repo)
		return c.JSON(resp.Status, resp)
	})

	r.POST("/:id/leave-submission", func(c echo.Context) error {
		admin := jwt.Authorize(c)
		if admin == nil {
			return nil
		}

		var req adapter.LeaveSubmissionReq
		_ = c.Bind(&req)
		if err := validator.Validate(req); err != nil {
			resp := err.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		timeout := 15 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)

		defer cancel()

		repo := EmployeeRepo{
			ctx: ctx,
			db:  sqlite.Db,
		}

		resp := adapter.HandleSubmitLeave(req, &repo)
		return c.JSON(resp.Status, resp)
	})
}
