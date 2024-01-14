package adapter

import errors "go-api-echo/internal/pkg/helpers/helpers_errors"

type AuthRepoInterface interface {
	GetUser(email string) (*User, *errors.Error)
	CreateUser(user *User) (*User, *errors.Error)
}
