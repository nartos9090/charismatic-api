package adapter

import errors "go-api-echo/internal/pkg/helpers/errors"

type AuthRepoInterface interface {
	GetAdminProfile(email string) (*Admin, *errors.Error)
}
