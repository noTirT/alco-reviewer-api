package util

import (
	"fmt"

	"github.com/go-playground/validator"
)

type ValidationError struct {
	validator.FieldError
}

func (v ValidationError) Error() string {
	if v.Tag() == "required" {
		return fmt.Sprintf("%s is required", v.Field())
	}

	return fmt.Sprintf(
		"key: '%s' Error: Field validation for '%s' failed on the '%s' tag",
		v.Namespace(),
		v.Field(),
		v.Tag())
}

type ValidationErrors []ValidationError

func (v ValidationErrors) Errors() []string {
	errs := []string{}
	for _, err := range v {
		errs = append(errs, err.Error())
	}
	return errs
}

type Validation struct {
	validate *validator.Validate
}

func NewValidation() *Validation {
	validate := validator.New()
	return &Validation{validate: validate}
}

func (v *Validation) Validate(i interface{}) ValidationErrors {
	errs := v.validate.Struct(i)
	if errs == nil {
		return nil
	}

	var returnErrs ValidationErrors
	for _, err := range errs.(validator.ValidationErrors) {
		ve := ValidationError{err}
		returnErrs = append(returnErrs, ve)
	}
	return returnErrs
}
