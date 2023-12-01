package jwt

import (
	errors "go-api-echo/internal/pkg/helpers/errors"
)

var RoleAdmin = "Admin"
var RoleAdminCode = "ADMIN"

var RoleUser = "User"
var RoleUserCode = "USER"

func AuthorizeAdmin(role string) *errors.Error {
	if role != RoleAdmin {
		errRes := *errors.UnauthorizedError
		errRes.AddError("can't access this feature")

		return &errRes
	}

	return nil
}
