package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Register(name string, validator ValidatorFilter) {
	registerValidator(name, validator)
}

func validateArray(context *ValidatorContext, normalizedValue *NormalizedValue, errors *Errors) {
	valueType := reflect.ValueOf(normalizedValue.Value)
	for i := 0; i < valueType.Len(); i++ {
		validateAny(context, valueType.Index(i).Interface(), errors)
	}
}

func validateMap(context *ValidatorContext, normalizedValue *NormalizedValue, errors *Errors) {
	valueType := reflect.ValueOf(normalizedValue.Value)
	for _, key := range valueType.MapKeys() {
		validateAny(context, valueType.MapIndex(key).Interface(), errors)
	}
}

func validateStruct(context *ValidatorContext, normalizedValue *NormalizedValue, errors *Errors) {
	for _, field := range getFields(normalizedValue.Value, "validate") {
		normalizedFieldValue, err := normalizeValue(field.Value, false)

		if err != nil {
			errors.Add(NewValidatorError(field.Name, "MISSING_TAG_NAME", strings.Replace(err.Error(), "{field}", field.Name, 1)))
			break
		}

		context.SetField(field)
		context.SetValue(normalizedFieldValue)

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

		validateAny(context, context.Value, errors)
	}
}

func validateAny(context *ValidatorContext, value interface{}, errors *Errors) {
	var normalizedValue *NormalizedValue

	if typedValue, ok := value.(*NormalizedValue); ok {
		normalizedValue = typedValue
	} else {
		normalizedValue, _ = normalizeValue(value, false)
	}

	switch normalizedValue.OriginalKind {
	case reflect.Array, reflect.Slice:
		validateArray(context, normalizedValue, errors)
	case reflect.Map:
		validateMap(context, normalizedValue, errors)
	case reflect.Struct:
		context.SetParent(normalizedValue.Value)
		validateStruct(context, normalizedValue, errors)
	}
}

func Validate(value interface{}) *Errors {
	errors := NewErrors()
	context := NewValidatorContext()

	validateAny(context, value, errors)

	return errors
}
