package errors

import "net/http"

var DuplicateEntryError = &Error{
	HttpCode: http.StatusConflict,
	Message:  `error: duplicate entry`,
	Errors:   []string{},
}

var BadRequestError = &Error{
	HttpCode: http.StatusBadRequest,
	Message:  `error: bad request`,
	Errors:   []string{},
}

var NotFoundError = &Error{
	HttpCode: http.StatusNotFound,
	Message:  `error: not found`,
	Errors:   []string{},
}

var UnauthorizedError = &Error{
	HttpCode: http.StatusUnauthorized,
	Message:  `error: unauthorized`,
	Errors:   []string{},
}

var ForbiddenError = &Error{
	HttpCode: http.StatusForbidden,
	Message:  `error: you don't have access'`,
	Errors:   []string{},
}

var InternalServerError = &Error{
	HttpCode: http.StatusInternalServerError,
	Message:  `error: internal server error`,
	Errors:   []string{},
}

var BadGatewayError = &Error{
	HttpCode: http.StatusBadGateway,
	Message:  `error: bad gateway`,
	Errors:   []string{},
}
