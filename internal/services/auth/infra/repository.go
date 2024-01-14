package infra

import (
	"context"
	"go-api-echo/internal/pkg/helpers/helpers_errors"
	"go-api-echo/internal/services/auth/adapter"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	ctx context.Context
	db  *sqlx.DB
}

func (r AuthRepo) GetUser(email string) (*adapter.User, *helpers_errors.Error) {
	admin := adapter.User{}
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
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("error getting login info")

		return nil, &sqlErr
	}

	return &admin, nil
}

func (r AuthRepo) CreateUser(user *adapter.User) (*adapter.User, *helpers_errors.Error) {
	_, err := r.db.ExecContext(
		r.ctx,
		`
		INSERT INTO user (
			email,
			fullname,
			passwd,
			passwdSalt,
			provider,
			provider_id
		) VALUES (?, ?, ?, ?, ?, ?)
		`,
		user.Email,
		user.FullName,
		user.Password,
		user.PasswordSalt,
		user.Provider,
		user.ProviderID,
	)
	if err != nil {
		sqlErr := helpers_errors.FromSql(err)
		sqlErr.AddError("error creating user")

		return nil, &sqlErr
	}

	return r.GetUser(user.Email)
}
