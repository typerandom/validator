package main

import (
	"errors"
	"strconv"
)

type ValidatorFn func(value interface{}, options string) error

func IsNotEmpty(value interface{}, options string) error {
	validateString := func(text *string) error {
		if text == nil || len(*text) == 0 {
			return errors.New("Value cannot be empty.")
		}
		return nil
	}

	validateInt := func(num *int) error {
		if num == nil {
			return errors.New("Value cannot be empty.")
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
	}

	return nil
}

func IsMin(value interface{}, options string) error {
	minValue, err := strconv.Atoi(options)

	if err != nil {
		return errors.New("Unable to parse 'min' validator value.")
	}

	validateString := func(text *string) error {
		if text == nil || len(*text) < minValue {
			return errors.New("Value is shorter than " + strconv.Itoa(minValue) + " characters.")
		}
		return nil
	}

	validateInt := func(num *int) error {
		if num == nil || *num > minValue {
			return errors.New("Value cannot be less than " + strconv.Itoa(minValue) + ".")
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

	return nil
}

func IsMax(value interface{}, options string) error {
	minValue, err := strconv.Atoi(options)

	if err != nil {
		return errors.New("Unable to parse 'max' validator value.")
	}

	validateString := func(text *string) error {
		if text != nil && len(*text) > minValue {
			return errors.New("Value is longer than " + strconv.Itoa(minValue) + " characters.")
		}
		return nil
	}

	validateInt := func(num *int) error {
		if num != nil && *num > minValue {
			return errors.New("Value cannot be greater than " + strconv.Itoa(minValue) + ".")
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

	return nil
}

func registerDefaultValidators() {
	registerValidator("not_empty", IsNotEmpty)
	registerValidator("min", IsMin)
	registerValidator("max", IsMax)
}
