package main

import (
	"errors"
	"reflect"
	"strconv"
	"unicode"
)

type UnsupportedTypeError struct {
	Type    reflect.Type
	Message string
}

func NewUnsupportedTypeError(validatorName string, value interface{}) *UnsupportedTypeError {
	valueType := reflect.TypeOf(value)
	return &UnsupportedTypeError{
		Type:    valueType,
		Message: "Validator with name '" + validatorName + "' on struct '{struct}' and field '{field}' is not supported.",
	}
}

func (this *UnsupportedTypeError) Error() string {
	return this.Message
}

type ValidatorContext struct {
	Value        interface{}
	StopValidate bool
}

func NewValidatorContext(value interface{}) *ValidatorContext {
	return &ValidatorContext{
		Value: value,
	}
}

type ValidatorFilter func(context *ValidatorContext, options []string) error

func IsEmpty(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'empty' does not support any arguments.")
	}

	validateString := func(text *string) error {
		if text == nil || len(*text) == 0 {
			context.StopValidate = true
		}
		return nil
	}

	validateInt := func(num *int64) error {
		if num == nil || *num == 0 {
			context.StopValidate = true
		}
		return nil
	}

	switch typedValue := context.Value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	case *int64:
		return validateInt(typedValue)
	case int64:
		return validateInt(&typedValue)
	default:
		if reflect.ValueOf(typedValue).IsNil() {
			context.StopValidate = true
			return nil
		}
	}

	return NewUnsupportedTypeError("empty", context.Value)
}

func IsNotEmpty(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'not_empty' does not support any arguments.")
	}

	validateString := func(text *string) error {
		if text == nil || len(*text) == 0 {
			return errors.New("{field} cannot be empty.")
		}
		return nil
	}

	validateInt := func(num *int64) error {
		if num == nil || *num == 0 {
			return errors.New("{field} cannot be empty.")
		}
		return nil
	}

	switch typedValue := context.Value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	case *int64:
		return validateInt(typedValue)
	case int64:
		return validateInt(&typedValue)
	default:
		valueType := reflect.ValueOf(context.Value)
		if valueType.IsNil() {
			return errors.New("{field} cannot be empty.")
		} else if valueType.Kind() == reflect.Ptr && valueType.Elem().Kind() == reflect.Struct {
			return nil
		}
	}

	return NewUnsupportedTypeError("not_empty", context.Value)
}

func IsMin(context *ValidatorContext, options []string) error {
	if len(options) != 1 {
		return errors.New("Validator 'min' requires a single argument.")
	}

	minValue, err := strconv.Atoi(options[0])

	if err != nil {
		return errors.New("Unable to parse 'min' validator value.")
	}

	validateString := func(text *string) error {
		if text == nil || len(*text) < minValue {
			return errors.New("{field} cannot be shorter than " + strconv.Itoa(minValue) + " characters.")
		}
		return nil
	}

	validateInt := func(num *int64) error {
		if num == nil || *num < int64(minValue) {
			return errors.New("{field} cannot be less than " + strconv.Itoa(minValue) + ".")
		}
		return nil
	}

	switch typedValue := context.Value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	case *int64:
		return validateInt(typedValue)
	case int64:
		return validateInt(&typedValue)
	}

	return NewUnsupportedTypeError("min", context.Value)
}

func IsMax(context *ValidatorContext, options []string) error {
	if len(options) != 1 {
		return errors.New("Validator 'max' requires a single argument.")
	}

	minValue, err := strconv.Atoi(options[0])

	if err != nil {
		return errors.New("Unable to parse 'max' validator value.")
	}

	validateString := func(text *string) error {
		if text != nil && len(*text) > minValue {
			return errors.New("{field} is longer than " + strconv.Itoa(minValue) + " characters.")
		}
		return nil
	}

	validateInt := func(num *int64) error {
		if num != nil && *num > int64(minValue) {
			return errors.New("{field} cannot be greater than " + strconv.Itoa(minValue) + ".")
		}
		return nil
	}

	switch typedValue := context.Value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	case *int64:
		return validateInt(typedValue)
	case int64:
		return validateInt(&typedValue)
	}

	return NewUnsupportedTypeError("max", context.Value)
}

func IsLowerCase(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'lowercase' does not support any arguments.")
	}

	validateString := func(text *string) error {
		if text == nil || len(*text) == 0 {
			return nil
		}

		for _, char := range *text {
			if unicode.IsLetter(char) && !unicode.IsLower(char) {
				return errors.New("{field} must be in lower case.")
			}
		}

		return nil
	}

	switch typedValue := context.Value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	}

	return NewUnsupportedTypeError("lowercase", context.Value)
}

func IsUpperCase(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'uppercase' does not support any arguments.")
	}

	validateString := func(text *string) error {
		if text == nil || len(*text) == 0 {
			return nil
		}

		for _, char := range *text {
			if unicode.IsLetter(char) && !unicode.IsUpper(char) {
				return errors.New("{field} must be in upper case.")
			}
		}

		return nil
	}

	switch typedValue := context.Value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	}

	return NewUnsupportedTypeError("uppercase", context.Value)
}

func IsNumeric(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'numeric' does not support any arguments.")
	}

	validateString := func(text *string) error {
		if text == nil || len(*text) == 0 {
			return errors.New("{field} must be numeric.")
		}

		value, err := strconv.ParseInt(*text, 10, 32)

		if err != nil {
			return errors.New("{field} must contain numbers only.")
		}

		context.Value = value

		return nil
	}

	switch typedValue := context.Value.(type) {
	case *string:
		return validateString(typedValue)
	case string:
		return validateString(&typedValue)
	}

	return NewUnsupportedTypeError("numeric", context.Value)
}

/*
IsHex
IsType
IsISO8601
IsUnixTime
IsEmail
IsUrl
IsFilePath
IsType				type(string)
IsInteger
IsDecimal
IsIP
IsRegexMatch
IsUUID
IsNumeric*/

func registerDefaultValidators() {
	registerValidator("empty", IsEmpty)
	registerValidator("not_empty", IsNotEmpty)
	registerValidator("min", IsMin)
	registerValidator("max", IsMax)
	registerValidator("lowercase", IsLowerCase)
	registerValidator("uppercase", IsUpperCase)
	registerValidator("numeric", IsNumeric)
}
