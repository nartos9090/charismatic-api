package auth_infra

import (
	"context"
	errors "go-api-echo/internal/pkg/helpers/errors"
	adapter "go-api-echo/internal/services/auth/adapter"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	ctx context.Context
	db  *sqlx.DB
	tx  *sqlx.Tx
}

func (r *AuthRepo) BeginTransaction() *errors.Error {
	tx, err := r.db.Beginx()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error starting transaction`)

		return &sqlErr
	}

	r.tx = tx

	return nil
}

func (r AuthRepo) CommitTransaction() *errors.Error {
	err := r.tx.Commit()
	if err != nil {
		sqlErr := errors.FromSql(err)
		sqlErr.AddError(`error committing transaction`)

		return &sqlErr
	}

	return nil
}

func (r AuthRepo) RollbackTransaction() {
	_ = r.tx.Rollback()
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
