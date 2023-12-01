package utils

import (
	"math"
	"reflect"

	errors "go-api-echo/internal/pkg/helpers/errors"
)

func RoundToNearestTenth(raw float64) (res float64) {
	res = math.Round(raw*10) / 10

	return
}

func RoundToNearestHundredth(raw float64) (res float64) {
	res = math.Round(raw*100) / 100

	return
}

func SetField[T any](v any, name string, value T) *errors.Error {
	var sample T
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr || rv.Elem().Kind() != reflect.Struct {
		err := *errors.InternalServerError
		err.AddError("v must be a pointer to struct")

		return &err
	}

	rv = rv.Elem()

	fv := rv.FieldByName(name)
	if !fv.IsValid() {
		err := *errors.InternalServerError
		err.AddErrorf("not a valid field name: %s", name)

		return &err
	}

	if !fv.CanSet() {
		err := *errors.InternalServerError
		err.AddErrorf("can't set field: %s", name)

		return &err
	}

	if fv.Kind() != reflect.ValueOf(sample).Kind() {
		err := *errors.InternalServerError
		err.AddErrorf("%s not a valid type", name)

		return &err
	}

	fv.Set(reflect.ValueOf(value))
	return nil
}
