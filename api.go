package main

import (
	"fmt"
)

func Register(name string, validator ValidatorFn) {
	registerValidator(name, validator)
}

func Validate(value interface{}) *Errors {
	var errors *Errors

	for _, field := range getTaggedFields(value, "validate") {
		for _, tag := range field.Tags {
			if validate, err := getValidator(tag.Name); err == nil {
				if err = validate(field.Value, tag.Options); err != nil {
					if errors == nil {
						errors = NewErrors()
					}
					errors.Add(NewValidatorError(field.Name, tag.Name, fmt.Sprintf(err.Error(), field.Name)))
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
