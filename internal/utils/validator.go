package utils

import (
	"github.com/go-playground/validator/v10"
)

// type ErrorResponse struct {
// 	FailedField string
// 	Tag         string
// 	Value       string
// }

var validate = validator.New()

func ValidateStruct(obj any) []string {
	// fmt.Println(&validate)
	var errors []string
	err := validate.Struct(obj)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element string
			element = "field: " + err.StructNamespace()
			element += " is " + err.Tag()
			element += " " + err.Param()
			errors = append(errors, element)
		}
	}
	return errors
}
