package main

import (
	"errors"
	"reflect"
	"strconv"
	"unicode"
)

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
IsUUID*/

type ValidatorContext struct {
	Errors       *Errors
	Source       interface{}
	Value        interface{}
	OriginalKind reflect.Kind
	Field        *reflectedField
	IsNil        bool
	StopValidate bool
}

func (this *ValidatorContext) SetValue(normalized *normalizedValue) {
	this.Value = normalized.Value
	this.OriginalKind = normalized.OriginalKind
	this.IsNil = normalized.IsNil
}

func (this *ValidatorContext) SetSource(source interface{}) {
	this.Source = source
}

func (this *ValidatorContext) SetField(field *reflectedField) {
	this.Field = field
}

type ValidatorFilter func(context *ValidatorContext, options []string) error

func emptyValidator(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'empty' does not support arguments.")
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
	case int64:
		if typedValue == 0 {
			context.StopValidate = true
		}
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

func notEmptyValidator(context *ValidatorContext, options []string) error {
	if len(options) > 0 {
		return errors.New("Validator 'not_empty' does not support arguments.")
	}

	if context.IsNil {
		return errors.New("{field} cannot be empty.")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if len(typedValue) == 0 {
			return errors.New("{field} cannot be empty.")
		}
	case int64:
		if typedValue == 0 {
			return errors.New("{field} cannot be empty.")
		}
	case float64:
		if typedValue == 0 {
			return errors.New("{field} cannot be empty.")
		}
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
	}

	return nil
}

func minValidator(context *ValidatorContext, options []string) error {
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

	return UnsupportedTypeError
}

func maxValidator(context *ValidatorContext, options []string) error {
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

	return UnsupportedTypeError
}

func lowerCaseValidator(context *ValidatorContext, options []string) error {
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

	return UnsupportedTypeError
}

func upperCaseValidator(context *ValidatorContext, options []string) error {
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

	return UnsupportedTypeError
}

func numericValidator(context *ValidatorContext, options []string) error {
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

	return UnsupportedTypeError
}

func funcValidator(context *ValidatorContext, options []string) error {
	var funcName string

	switch len(options) {
	case 0:
		funcName = "Validate" + context.Field.Name
	case 1:
		funcName = options[0]
	default:
		return errors.New("Validator 'func' does not support more than 1 argument.")
	}

	returnValues, err := callMethod(context.Source, funcName, context)

	if err != nil {
		if err == InvalidMethodError {
			return errors.New("Validation method '" + context.Field.Parent.FullName(funcName) + "' does not exist.")
		}
		return err
	}

	if len(returnValues) == 1 {
		if returnValues[0] == nil {
			return nil
		} else if err, ok := returnValues[0].(error); ok {
			return errors.New(err.Error())
		} else {
			return errors.New("Invalid return value of validation method '" + context.Field.Parent.FullName(funcName) + "'. Return value must be of type 'error'.")
		}
	}

	return UnsupportedTypeError
}

func registerDefaultValidators() {
	registerValidator("empty", emptyValidator)
	registerValidator("not_empty", notEmptyValidator)
	registerValidator("min", minValidator)
	registerValidator("max", maxValidator)
	registerValidator("lowercase", lowerCaseValidator)
	registerValidator("uppercase", upperCaseValidator)
	registerValidator("numeric", numericValidator)
	registerValidator("func", funcValidator)
}
