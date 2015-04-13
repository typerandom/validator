package main

import (
	"errors"
	"reflect"
	"strconv"
)

type UnsupportedValueTypeError struct {
	Type    reflect.Type
	Message string
}

func NewUnsupportedValueTypeError(value interface{}) *UnsupportedValueTypeError {
	valueType := reflect.TypeOf(value)
	return &UnsupportedValueTypeError{
		Type:    valueType,
		Message: "Type '" + valueType.Name() + "' on field '%s' is not supported.",
	}
}

func (this *UnsupportedValueTypeError) Error() string {
	return this.Message
}

type ValidatorFn func(value interface{}, options string) error

func IsNotEmpty(value interface{}, options string) error {
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

	return NewUnsupportedValueTypeError(value)
}

func IsMin(value interface{}, options string) error {
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
		if num == nil || *num > minValue {
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

	return NewUnsupportedValueTypeError(value)
}

func IsMax(value interface{}, options string) error {
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

	return NewUnsupportedValueTypeError(value)
}

func registerDefaultValidators() {
	registerValidator("not_empty", IsNotEmpty)
	registerValidator("min", IsMin)
	registerValidator("max", IsMax)
}
