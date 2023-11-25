package auth_adapter

import errors "go-api-echo/internal/pkg/helpers/errors"

type AuthRepoInterface interface {
	BeginTransaction() *errors.Error
	CommitTransaction() *errors.Error
	RollbackTransaction()

	GetAdminProfile(email string) (*Admin, *errors.Error)
}
