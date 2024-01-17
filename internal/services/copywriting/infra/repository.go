package infra

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-api-echo/config"
	errors "go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/copywriting/adapter"
	"go-api-echo/internal/services/copywriting/entity"
)

type CopywritingProjectRepository struct {
	ctx context.Context
	db  *sqlx.DB
}

func (r CopywritingProjectRepository) Create(userID int, req adapter.CreateCopywritingRepoReq) (*adapter.CopywritingProjectDetail, *errors.Error) {
	res, err := r.db.ExecContext(
		r.ctx,
		`INSERT INTO copywriting_project (
			user_id,
			 title,
			 product_image,
			 brand_name,
			 market_target,
			 superiority,
			 result
		 ) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		userID,
		req.Title,
		req.ProductImage,
		req.BrandName,
		req.MarketTarget,
		req.Superiority,
		req.Result,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error when inserting data")
		return nil, &sqlErr
	}

	id, err := res.LastInsertId()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error when getting last inserted id")
		return nil, &sqlErr
	}

	project, errs := r.GetDetail(int(id), userID)
	if errs != nil {
		return nil, errs
	}

	return project, nil
}
func (r CopywritingProjectRepository) GetLists(userID int) (*[]entity.Copywriting, *errors.Error) {
	rows, err := r.db.QueryxContext(
		r.ctx,
		`SELECT 
			id,
			user_id, 
			title, 
			product_image,
			brand_name, 
			market_target, 
			superiority 
		FROM copywriting_project
		WHERE user_id = ?;`,
		userID,
	)

	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError("error when querying data")
		return nil, &sqlErr
	}

	projects := []entity.Copywriting{}
	for rows.Next() {
		var project entity.Copywriting
		err := rows.StructScan(&project)
		if err != nil {
			sqlErr := errors.FromSql(err)
			sqlErr.AddError("error when scanning data")
			return nil, &sqlErr
		}
		project.ProductImage = config.GlobalEnv.BaseURL + project.ProductImage
		projects = append(projects, project)
	}

	return &projects, nil
}
func (r CopywritingProjectRepository) GetDetail(projectID int, userID int) (*adapter.CopywritingProjectDetail, *errors.Error) {
	var project adapter.CopywritingProjectDetail
	err := r.db.QueryRowxContext(
		r.ctx,
		`SELECT 
			id,
			user_id, 
			title, 
			product_image,
			brand_name, 
			market_target, 
			superiority,
			result
		FROM copywriting_project
		WHERE id = ? AND user_id = ?;`,
		projectID,
		userID,
	).StructScan(&project)

	if err != nil {
		sqlErr := errors.FromSql(err)
		return nil, &sqlErr
	}

	project.ProductImage = config.GlobalEnv.BaseURL + project.ProductImage

	return &project, nil
}
