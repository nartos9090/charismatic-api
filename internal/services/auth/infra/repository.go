package infra

import (
	"context"
	"go-api-echo/internal/pkg/helpers/errors"
	"go-api-echo/internal/services/auth/adapter"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	ctx context.Context
	db  *sqlx.DB
}

func (r AuthRepo) GetAdminProfile(email string) (*adapter.Admin, *errors.Error) {
	admin := adapter.Admin{}

	err := r.db.GetContext(
		r.ctx,
		&admin,
		`
		SELECT
			id,
			email,
			fullname,
			passwd,
			passwdSalt
		FROM admins
		WHERE email = ?
		`,
		email,
	)
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error getting login info`)

		return nil, &sqlErr
	}

	return &admin, nil
}
