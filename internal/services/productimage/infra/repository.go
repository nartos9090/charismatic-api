package infra

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/productimage/adapter"
	"go-api-echo/internal/services/productimage/entity"
)

type ProductImageRepository struct {
	ctx context.Context
	db  *sqlx.DB
}

func (r ProductImageRepository) CreateProductImage(userID int, product adapter.CreateProductImageRepoReq) (int, *helpers_errors.Error) {
	res, err := r.db.ExecContext(
		r.ctx,
		`INSERT INTO product_image (
			user_id,
		    title,
		    image_url,
		    mask_url
	    ) VALUES (?, ?, ?, ?)`,
		userID,
		product.Title,
		product.ImageUrl,
		product.MaskUrl,
	)
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("failed to create product image")
		return 0, &sqlErr
	}

	productID, err := res.LastInsertId()
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("failed to create product image")
		return 0, &sqlErr
	}

	return int(productID), nil
}
func (r ProductImageRepository) CreateGeneratedProductImage(userID int, generated adapter.AddGeneratedProductImageRepoReq) (*entity.ProductImageGenerated, *helpers_errors.Error) {
	res, err := r.db.ExecContext(
		r.ctx,
		"INSERT INTO product_image_edited (product_image_id, image_url, prompt) VALUES (?, ?, ?)",
		generated.ProductImageID,
		generated.ImageUrl,
		generated.Prompt,
	)
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("failed to create generated product image")
		return nil, &sqlErr
	}

	generatedID, err := res.LastInsertId()
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("failed to create generated product image")
		return nil, &sqlErr
	}

	return r.GetGeneratedProductImage(userID, int(generatedID))
}

func (r ProductImageRepository) GetGeneratedProductImage(userID, id int) (*entity.ProductImageGenerated, *helpers_errors.Error) {
	var generated entity.ProductImageGenerated
	row := r.db.QueryRowx(
		`SELECT
    				pie.id,
    				pie.product_image_id,
    				pie.image_url,
    				pie.prompt
			FROM product_image_edited pie
			LEFT JOIN charismatic_dev.product_image pi on pie.product_image_id = pi.id
			WHERE pie.id = ?
				AND pi.user_id = ?`,
		id,
		userID,
	)

	err := row.StructScan(&generated)
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("failed to get generated product image")
		return nil, &sqlErr
	}

	return &generated, nil
}

func (r ProductImageRepository) GetGeneratedProductImageList(productID int) (*[]entity.ProductImageGenerated, *helpers_errors.Error) {
	var generated []entity.ProductImageGenerated
	rows, err := r.db.Queryx(
		`SELECT
					id,
					product_image_id,
					image_url,
					prompt
			FROM product_image_edited
			WHERE product_image_id = ?`,
		productID,
	)
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("failed to get generated product image list")
		return nil, &sqlErr
	}

	for rows.Next() {
		var product entity.ProductImageGenerated
		err := rows.StructScan(&product)
		if err != nil {
			sqlErr := helpers_errors.FromSql(err)
			sqlErr.AddError("failed to scan generated product image")
			return nil, &sqlErr
		}
		generated = append(generated, product)
	}

	return &generated, nil

}

func (r ProductImageRepository) GetProductImageList(userID int) (*[]entity.ProductImage, *helpers_errors.Error) {
	products := []entity.ProductImage{}
	rows, err := r.db.Queryx(
		`SELECT
					id,
					user_id,
					title,
					image_url,
					mask_url
			FROM product_image
			WHERE user_id = ?`,
		userID,
	)
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("failed to get product image list")
		return nil, &sqlErr
	}

	for rows.Next() {
		var product entity.ProductImage
		err := rows.StructScan(&product)
		if err != nil {
			sqlErr := helpers_errors.FromSql(err)
			sqlErr.AddError("failed to scan product image")
			return nil, &sqlErr
		}
		products = append(products, product)
	}

	return &products, nil

}
func (r ProductImageRepository) GetProductImage(userID, productID int) (*adapter.ProductImageDetail, *helpers_errors.Error) {
	var product adapter.ProductImageDetail
	row := r.db.QueryRowx(
		`SELECT
					id,
					user_id,
					title,
					image_url,
					mask_url
			FROM product_image
			WHERE id = ? AND user_id = ?`,
		productID,
		userID,
	)

	err := row.StructScan(&product.ProductImage)
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("failed to get product image")
		return nil, &sqlErr
	}

	generatedImages, errs := r.GetGeneratedProductImageList(productID)
	if err != nil {
		return nil, errs
	}

	product.GeneratedImages = *generatedImages

	return &product, nil
}
