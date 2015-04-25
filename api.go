package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Register(name string, validator ValidatorFilter) {
	registerValidator(name, validator)
}

func validateField(field *taggedField, normalizedValue *NormalizedValue, errors *Errors) {
	context := NewValidatorContext(normalizedValue)

	for _, tag := range field.Tags {
		if validate, err := getValidator(tag.Name); err == nil {
			if err = validate(context, tag.Options); err != nil {
				errors.Add(NewValidatorError(field.Name, tag.Name, strings.Replace(err.Error(), "{field}", field.Name, 1)))
			}
			if context.StopValidate {
				break
			}
		} else {
			errors.Add(NewValidatorError(field.Name, tag.Name, fmt.Sprintf("Validator '%s' used on field '%s' does not exist.", tag.Name, field.Name)))
			break
		}
	}

	validateAny(normalizedValue, errors)
}

func validateArray(normalizedValue *NormalizedValue, errors *Errors) {
	valueType := reflect.ValueOf(normalizedValue.Value)
	for i := 0; i < valueType.Len(); i++ {
		validateAny(valueType.Index(i).Interface(), errors)
	}
}

func validateStruct(normalizedValue *NormalizedValue, errors *Errors) {
	for _, field := range getTaggedFields(normalizedValue.Value, "validate") {

		normalizedFieldValue, err := normalizeValue(field.Value, false)

		if err != nil {
			errors.Add(NewValidatorError(field.Name, "MISSING_TAG_NAME", strings.Replace(err.Error(), "{field}", field.Name, 1)))
			break
		}

		validateField(field, normalizedFieldValue, errors)
	}
}

func validateAny(value interface{}, errors *Errors) {
	var normalizedValue *NormalizedValue

	if typedValue, ok := value.(*NormalizedValue); ok {
		normalizedValue = typedValue
	} else {
		normalizedValue, _ = normalizeValue(value, false)
	}

	switch normalizedValue.OriginalKind {
	case reflect.Array, reflect.Slice:
		validateArray(normalizedValue, errors)
	case reflect.Struct:
		validateStruct(normalizedValue, errors)
	}
}

func Validate(value interface{}) *Errors {
	errors := NewErrors()

	validateAny(value, errors)

	return errors
}
