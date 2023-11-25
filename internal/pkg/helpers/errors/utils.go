package errors

import (
	"fmt"
	response "go-api-echo/internal/pkg/helpers/response"
)

func (err *Error) ToHttpRes() response.HttpRes {
	return response.HttpRes{
		Status:  err.HttpCode,
		Message: err.Message,
		Data:    nil,
		Errors:  err.Errors,
	}
}

func (err *Error) AddError(a ...any) {
	err.Errors = append(err.Errors, fmt.Sprint(a...))
}

func (err *Error) AddErrorf(format string, a ...any) {
	err.Errors = append(err.Errors, fmt.Sprintf(format, a...))
}
