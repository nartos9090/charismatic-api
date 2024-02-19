package infra

import (
	"bytes"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go-api-echo/external/services/clipdrop_service"
	"go-api-echo/internal/pkg/db/mysql"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/pkg/jwt"
	"go-api-echo/internal/services/productimage/adapter"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ProductImageRoute(g *echo.Group) {
	r := g.Group("/product-image")
	defer fmt.Printf(":: route /product-image created\n")

	r.POST("/create-sync", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		var req adapter.CreateProductImageReq
		_ = c.Bind(&req)

		imageFile, err := c.FormFile("image")
		if err != nil {
			comErr := *helpers_errors.BadRequestError
			comErr.AddError("invalid image")
			resp := comErr.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		// Open the imageFile
		src, err := imageFile.Open()
		if err != nil {
			comErr := *helpers_errors.BadRequestError
			comErr.AddError("invalid image")
			resp := comErr.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}
		defer src.Close()

		// Create the destination imageFile
		uploadPath := "./public/images/" // Change this path as needed
		if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
			return err
		}

		imageFilename := filepath.Join(uploadPath, strconv.Itoa(int(time.Now().Unix()))+"-"+imageFile.Filename)
		dst, err := os.Create(imageFilename)
		if err != nil {
			comErr := *helpers_errors.BadRequestError
			comErr.AddError("invalid image")
			resp := comErr.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		if strings.HasSuffix(imageFilename, ".jpg") || strings.HasSuffix(imageFilename, ".jpeg") {
			imgBytes, err := os.ReadFile(imageFilename)
			if err != nil {
				comErr := *helpers_errors.BadRequestError
				comErr.AddError("invalid image")
				resp := comErr.ToHttpRes()
				return c.JSON(resp.Status, resp)
			}

			imageFilename = filepath.Join(uploadPath, strconv.Itoa(int(time.Now().Unix()))+"-"+imageFile.Filename)
			img, err := jpeg.Decode(bytes.NewReader(imgBytes))
			if err != nil {
				comErr := *helpers_errors.BadRequestError
				comErr.AddError("invalid image")
				resp := comErr.ToHttpRes()
				return c.JSON(resp.Status, resp)
			}

			buf := new(bytes.Buffer)
			err = png.Encode(buf, img)
			if err != nil {
				comErr := *helpers_errors.BadRequestError
				comErr.AddError("invalid image")
				resp := comErr.ToHttpRes()
				return c.JSON(resp.Status, resp)
			}

			dst, err = os.Create(imageFilename)
			if err != nil {
				comErr := *helpers_errors.BadRequestError
				comErr.AddError("invalid image")
				resp := comErr.ToHttpRes()
				return c.JSON(resp.Status, resp)
			}

			if _, err = io.Copy(dst, buf); err != nil {
				comErr := *helpers_errors.BadRequestError
				comErr.AddError("invalid image")
				resp := comErr.ToHttpRes()
				return c.JSON(resp.Status, resp)
			}
		}

		req.BaseImage = imageFilename

		maskFile, err := c.FormFile("mask")
		if maskFile != nil {
			if err != nil {
				comErr := *helpers_errors.BadRequestError
				comErr.AddError("invalid mask")
				resp := comErr.ToHttpRes()
				return c.JSON(resp.Status, resp)
			}

			// Open the maskFile
			src, err = maskFile.Open()
			if err != nil {
				comErr := *helpers_errors.BadRequestError
				comErr.AddError("invalid mask")
				resp := comErr.ToHttpRes()
				return c.JSON(resp.Status, resp)
			}

			maskFilename := filepath.Join(uploadPath, strconv.Itoa(int(time.Now().Unix()))+"-"+maskFile.Filename+".png")
			dst, err = os.Create(maskFilename)
			if err != nil {
				comErr := *helpers_errors.BadRequestError
				comErr.AddError("invalid mask")
				resp := comErr.ToHttpRes()
				return c.JSON(resp.Status, resp)
			}

			if _, err = io.Copy(dst, src); err != nil {
				return err
			}

			defer dst.Close()

			req.MaskImage = maskFilename
		}

		repo := ProductImageRepository{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		clipdropRepo := clipdrop_service.CripdropServiceInterface{}

		resp := adapter.CreateProductImage(user.ID, req, repo, clipdropRepo, clipdropRepo)
		return c.JSON(resp.Status, resp)
	})

	r.GET("/list", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		repo := ProductImageRepository{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.GetProductImageList(user.ID, repo)

		return c.JSON(resp.Status, resp)
	})

	r.GET("/detail/:id", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			comErr := *helpers_errors.BadRequestError
			comErr.AddError("invalid id")
			resp := comErr.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		repo := ProductImageRepository{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.GetProductImage(user.ID, id, repo)

		return c.JSON(resp.Status, resp)
	})

	r.POST("/detail/:id/create-sync", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		var req adapter.GenerateBackgroundReq
		_ = c.Bind(&req)

		repo := ProductImageRepository{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		backgroundReplacerRepo := clipdrop_service.CripdropServiceInterface{}

		resp := adapter.GenerateImageBackground(user.ID, req, repo, backgroundReplacerRepo)

		return c.JSON(resp.Status, resp)
	})

	r.GET("/generated/:id", func(c echo.Context) error {
		user := jwt.Authorize(c)
		if user == nil {
			return nil
		}

		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			comErr := *helpers_errors.BadRequestError
			comErr.AddError("invalid id")
			resp := comErr.ToHttpRes()
			return c.JSON(resp.Status, resp)
		}

		repo := ProductImageRepository{
			ctx: context.Background(),
			db:  db_mysql.Db,
		}

		resp := adapter.GetGeneratedProductImage(user.ID, id, repo)

		return c.JSON(resp.Status, resp)
	})
}
