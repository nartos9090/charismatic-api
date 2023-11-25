package validator

import (
	"fmt"
	errors "go-api-echo/internal/pkg/helpers/errors"

	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	FailedField string `json:"failed_field"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateStruct(typeStruct interface{}) []*ErrorResponse {
	var res []*ErrorResponse

	if err := validate.Struct(typeStruct); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.Field()
			element.Tag = err.Tag()
			element.Value = err.Param()
			res = append(res, &element)
		}
	}

	return res
}

func Validate(typeStruct interface{}) *errors.Error {
	err := ValidateStruct(typeStruct)
	if err != nil {
		errRes := *errors.BadRequestError
		for _, v := range err {
			errMsg := fmt.Sprintf(`field %s must be type of %s, invalid value of %s`, v.FailedField, v.Tag, v.Value)
			errRes.Errors = append(errRes.Errors, errMsg)
		}
		return &errRes
	}

	return nil
}
