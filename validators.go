package main

import (
	"errors"
	"reflect"
	"strconv"
)

type UnsupportedTypeError struct {
	Type    reflect.Type
	Message string
}

func NewUnsupportedTypeError(value interface{}) *UnsupportedTypeError {
	valueType := reflect.TypeOf(value)
	return &UnsupportedTypeError{
		Type:    valueType,
		Message: "Validator with name '" + valueType.Name() + "' on struct '{StructName}' and field '{FieldName}' is not supported.",
	}
}

func (this *UnsupportedTypeError) Error() string {
	return this.Message
}

type ValidatorContext struct {
	StopValidate bool
}

func NewValidatorContext() *ValidatorContext {
	return &ValidatorContext{}
}

type ValidatorFn func(context *ValidatorContext, value interface{}, options string) error

func IsEmpty(context *ValidatorContext, value interface{}, options string) error {
	validateString := func(text *string) error {
		if text == nil || len(*text) == 0 {
			context.StopValidate = true
		}
		return nil
	}

	validateInt := func(num *int) error {
		if num == nil || *num == 0 {
			context.StopValidate = true
		}
		return nil
	}

	switch typedValue := value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	case *int:
		return validateInt(typedValue)
	case int:
		return validateInt(&typedValue)
	}

	return NewUnsupportedTypeError(value)
}

func IsNotEmpty(context *ValidatorContext, value interface{}, options string) error {
	validateString := func(text *string) error {
		if text == nil || len(*text) == 0 {
			return errors.New("%s cannot be empty.")
		}
		return nil
	}

	validateInt := func(num *int) error {
		if num == nil || *num == 0 {
			return errors.New("%s cannot be empty.")
		}
		return nil
	}

	switch typedValue := value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	case *int:
		return validateInt(typedValue)
	case int:
		return validateInt(&typedValue)
	}

	return NewUnsupportedTypeError(value)
}

func IsMin(context *ValidatorContext, value interface{}, options string) error {
	minValue, err := strconv.Atoi(options)

	if err != nil {
		return errors.New("Unable to parse 'min' validator value.")
	}

	validateString := func(text *string) error {
		if text == nil || len(*text) < minValue {
			return errors.New("%s cannot be shorter than " + strconv.Itoa(minValue) + " characters.")
		}
		return nil
	}

	validateInt := func(num *int) error {
		if num == nil || *num < minValue {
			return errors.New("%s cannot be less than " + strconv.Itoa(minValue) + ".")
		}
		return nil
	}

	switch typedValue := value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	case *int:
		return validateInt(typedValue)
	case int:
		return validateInt(&typedValue)
	}

	return NewUnsupportedTypeError(value)
}

func IsMax(context *ValidatorContext, value interface{}, options string) error {
	minValue, err := strconv.Atoi(options)

	if err != nil {
		return errors.New("Unable to parse 'max' validator value.")
	}

	validateString := func(text *string) error {
		if text != nil && len(*text) > minValue {
			return errors.New("%s is longer than " + strconv.Itoa(minValue) + " characters.")
		}
		return nil
	}

	validateInt := func(num *int) error {
		if num != nil && *num > minValue {
			return errors.New("%s cannot be greater than " + strconv.Itoa(minValue) + ".")
		}
		return nil
	}

	switch typedValue := value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	case *int:
		return validateInt(typedValue)
	case int:
		return validateInt(&typedValue)
	}

	return NewUnsupportedTypeError(value)
}

func registerDefaultValidators() {
	registerValidator("empty", IsEmpty)
	registerValidator("not_empty", IsNotEmpty)
	registerValidator("min", IsMin)
	registerValidator("max", IsMax)
}
