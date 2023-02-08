package utils

import "github.com/go-playground/validator/v10"

type ValidationErr struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ParseValidationError(err error) interface{} {
	var errors []*ValidationErr

	for _, err := range err.(validator.ValidationErrors) {
		var item ValidationErr
		item.Field = err.Field()
		item.Message = err.Field() + " " + err.Tag()
		errors = append(errors, &item)
	}
	return errors
}
