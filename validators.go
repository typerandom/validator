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
	Parent       interface{}
	Value        interface{}
	OriginalKind reflect.Kind
	Field        *reflectedField
	IsNil        bool
	StopValidate bool
}

func NewValidatorContext(parent interface{}, normalizedValue *NormalizedValue, field *reflectedField) *ValidatorContext {
	return &ValidatorContext{
		Parent:       parent,
		Value:        normalizedValue.Value,
		OriginalKind: normalizedValue.OriginalKind,
		Field:        field,
		IsNil:        normalizedValue.IsNil,
	}
}

type ValidatorFilter func(context *ValidatorContext, options []string) error

func EmptyValidator(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'empty' does not support any arguments.")
	}

	if context.IsNil {
		context.StopValidate = true
		return nil
	}

	switch typedValue := context.Value.(type) {
	case string:
		if len(typedValue) == 0 {
			context.StopValidate = true
		}
		return nil
	case int64:
		if typedValue == 0 {
			context.StopValidate = true
		}
		return nil
	}

	switch context.OriginalKind {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value).Len() == 0 {
			context.StopValidate = true
		}
	case reflect.Map:
		if len(reflect.ValueOf(context.Value).MapKeys()) == 0 {
			context.StopValidate = true
		}
	}

	return nil
}

func NotEmptyValidator(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'not_empty' does not support any arguments.")
	}

	if context.IsNil {
		return errors.New("{field} cannot be empty.")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if len(typedValue) == 0 {
			return errors.New("{field} cannot be empty.")
		}
		return nil
	case int64:
		if typedValue == 0 {
			return errors.New("{field} cannot be empty.")
		}
		return nil
	case float64:
		if typedValue == 0 {
			return errors.New("{field} cannot be empty.")
		}
		return nil
	}

	switch context.OriginalKind {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value).Len() == 0 {
			return errors.New("{field} cannot be empty.")
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value).MapKeys()) == 0 {
			return errors.New("{field} cannot be empty.")
		}
		return nil
	}

	return nil
}

func MinValidator(context *ValidatorContext, options []string) error {
	if len(options) != 1 {
		return errors.New("Validator 'min' requires a single argument.")
	}

	minValue, err := strconv.Atoi(options[0])

	if err != nil {
		return errors.New("Unable to parse 'min' validator value.")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if context.IsNil || len(typedValue) < minValue {
			return errors.New("{field} cannot be shorter than " + strconv.Itoa(minValue) + " characters.")
		}
		return nil
	case int64:
		if context.IsNil || typedValue < int64(minValue) {
			return errors.New("{field} cannot be less than " + strconv.Itoa(minValue) + ".")
		}
		return nil
	case float64:
		if context.IsNil || typedValue < float64(minValue) {
			return errors.New("{field} cannot be less than " + strconv.Itoa(minValue) + ".")
		}
		return nil
	}

	switch context.OriginalKind {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value).Len() < minValue {
			return errors.New("{field} cannot contain less than " + strconv.Itoa(minValue) + " items.")
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value).MapKeys()) < minValue {
			return errors.New("{field} cannot contain less than " + strconv.Itoa(minValue) + " keys.")
		}
		return nil
	}

	return NewUnsupportedTypeError("min", context.Value)
}

func MaxValidator(context *ValidatorContext, options []string) error {
	if len(options) != 1 {
		return errors.New("Validator 'max' requires a single argument.")
	}

	maxValue, err := strconv.Atoi(options[0])

	if err != nil {
		return errors.New("Unable to parse 'max' validator value.")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if !context.IsNil && len(typedValue) > maxValue {
			return errors.New("{field} is longer than " + strconv.Itoa(maxValue) + " characters.")
		}
		return nil
	case int64:
		if !context.IsNil && typedValue > int64(maxValue) {
			return errors.New("{field} cannot be greater than " + strconv.Itoa(maxValue) + ".")
		}
		return nil
	case float64:
		if !context.IsNil && typedValue > float64(maxValue) {
			return errors.New("{field} cannot be greater than " + strconv.Itoa(maxValue) + ".")
		}
		return nil
	}

	switch context.OriginalKind {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value).Len() > maxValue {
			return errors.New("{field} cannot contain more than " + strconv.Itoa(maxValue) + " items.")
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value).MapKeys()) > maxValue {
			return errors.New("{field} cannot contain more than " + strconv.Itoa(maxValue) + " keys.")
		}
		return nil
	}

	return NewUnsupportedTypeError("max", context.Value)
}

func LowerCaseValidator(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'lowercase' does not support any arguments.")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if context.IsNil || len(typedValue) == 0 {
			return nil
		}

		for _, char := range typedValue {
			if unicode.IsLetter(char) && !unicode.IsLower(char) {
				return errors.New("{field} must be in lower case.")
			}
		}

		return nil
	}

	return NewUnsupportedTypeError("lowercase", context.Value)
}

func UpperCaseValidator(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'uppercase' does not support any arguments.")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if context.IsNil || len(typedValue) == 0 {
			return nil
		}

		for _, char := range typedValue {
			if unicode.IsLetter(char) && !unicode.IsUpper(char) {
				return errors.New("{field} must be in upper case.")
			}
		}

		return nil
	}

	return NewUnsupportedTypeError("uppercase", context.Value)
}

func NumericValidator(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'numeric' does not support any arguments.")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if context.IsNil || len(typedValue) == 0 {
			return errors.New("{field} must be numeric.")
		}

		value, err := strconv.ParseInt(typedValue, 10, 32)

		if err != nil {
			return errors.New("{field} must contain numbers only.")
		}

		context.Value = value

		return nil
	case int64:
		return nil
	case float64:
		return nil
	}

	return NewUnsupportedTypeError("numeric", context.Value)
}

func FuncValidator(context *ValidatorContext, options []string) error {
	var funcName string

	switch len(options) {
	case 0:
		funcName = "Validate" + context.Field.Name
	case 1:
		funcName = options[0]
	default:
		return errors.New("Validator 'func' does not support more than 1 argument.")
	}

	returnValues, err := CallMethod(context.Parent, funcName, context.Value)

	if err != nil {
		if err == InvalidMethodError {
			return errors.New("Validation method '" + funcName + "' does not exist.")
		}
		return err
	}

	if len(returnValues) == 1 {
		if returnValues[0] == nil {
			return nil
		} else if err, ok := returnValues[0].(error); ok {
			return err
		} else {
			return errors.New("Invalid return value type. Must be error.")
		}
	}

	return errors.New("Invalid number of return values.")
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
	registerValidator("empty", EmptyValidator)
	registerValidator("not_empty", NotEmptyValidator)
	registerValidator("min", MinValidator)
	registerValidator("max", MaxValidator)
	registerValidator("lowercase", LowerCaseValidator)
	registerValidator("uppercase", UpperCaseValidator)
	registerValidator("numeric", NumericValidator)
	registerValidator("func", FuncValidator)
}
