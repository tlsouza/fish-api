package logic

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

func GetRequiredFieldError[T any](err error, req *T) error {

	if validationErrors, ok := err.(validator.ValidationErrors); ok {

		for _, fieldErr := range validationErrors {

			field, _ := reflect.TypeOf(req).Elem().FieldByName(fieldErr.Field())
			jsonField := GetJSONFieldName(field)

			if fieldErr.Tag() == "required" {
				return fmt.Errorf("the '%s' field is required", jsonField)
			}
		}
	}
	return err
}
