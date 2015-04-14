package main

import (
	"fmt"
	"strings"
)

func Register(name string, validator ValidatorFilter) {
	registerValidator(name, validator)
}

func Validate(value interface{}) *Errors {
	var errors *Errors

	for _, field := range getTaggedFields(value, "validate") {
		context := NewValidatorContext(field.Value)

		for _, tag := range field.Tags {
			if validate, err := getValidator(tag.Name); err == nil {
				if err = validate(context, tag.Options); err != nil {
					if errors == nil {
						errors = NewErrors()
					}
					errors.Add(NewValidatorError(field.Name, tag.Name, strings.Replace(err.Error(), "{field}", field.Name, 1)))
				}
				if context.StopValidate {
					break
				}
			} else {
				errors = NewErrors()
				errors.Add(NewValidatorError(field.Name, tag.Name, fmt.Sprintf("Validator '%s' used on field '%s' does not exist.", tag.Name, field.Name)))
				return errors
			}
		}
	}

	return errors
}
