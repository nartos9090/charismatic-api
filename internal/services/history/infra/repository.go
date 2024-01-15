package infra

import (
	"context"
	"github.com/jmoiron/sqlx"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/history/entity"
)

type HistoryRepoInterface struct {
	ctx context.Context
	db  *sqlx.DB
}

func (r HistoryRepoInterface) GetHistory(userID int) (*[]entity.History, *helpers_errors.Error) {
	histories := []entity.History{}

	rows, err := r.db.Queryx(`
		SELECT id, title as title, 'copywriting' as type, created_at
		FROM copywriting_project WHERE user_id = ?
		UNION
		SELECT id, product_title as title, 'video' as type, created_at
		FROM video_project WHERE user_id = ?
		UNION
		SELECT id, title as title, 'product_image' as type, created_at
		FROM product_image WHERE user_id = ?
		ORDER BY created_at DESC;
	`,
		userID,
		userID,
		userID,
	)

	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError(`Error when get history`)
		return nil, &sqlErr
	}

	for rows.Next() {
		var history entity.History
		err := rows.StructScan(&history)
		if err != nil {
			sqlErr := helpers_errors.FromSql(err)
			sqlErr.AddError(`Error when scanning history`)
			return nil, &sqlErr
		}
		histories = append(histories, history)
	}

	return &histories, nil
}
