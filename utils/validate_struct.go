package utils

import (
	"gochiapp/model"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func Validate[T any](t T) {

	validate = validator.New()
	err := validate.Struct(&t)

	if err != nil {

		var message []map[string]interface{}
		for _, e := range err.(validator.ValidationErrors) {

			message = append(message, map[string]interface{}{
				"field":   e.Field(),
				"message": customMessage(e),
			})

		}

		panic(model.ResponseFailWeb{
			Status:     "Failed",
			Error:      message,
			StatusCode: 400,
		})
	}
}

func customMessage(e validator.FieldError) string {

	switch e.Tag() {
	case "required":
		return "This field is required" + e.Param()
	case "min":
		return "length of this field should be more than " + e.Param() + " characters"
	case "max":
		return "length of this field should be less than " + e.Param() + " characters"
	case "alpha":
		return "Must be alphabet only"
	}

	return ""
}
