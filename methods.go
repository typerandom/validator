package main

import (
	"errors"
	"strconv"
)

type ValidatorStrategy func(value interface{}, options string) error

func IsEmpty(value interface{}, options string) error {
	switch typedValue := value.(type) {
	case string:
		if len(typedValue) > 0 {
			return errors.New("Value is not empty.")
		}
	}
	return nil
}

func IsMin(value interface{}, options string) error {
	minValue, _ := strconv.Atoi(options)
	switch typedValue := value.(type) {
	case *string:
		if len(*typedValue) < minValue {
			return errors.New("Value is shorter than " + strconv.Itoa(minValue) + " characters.")
		}
	}
	return nil
}

func IsMax(value interface{}, options string) error {
	return nil
}

var validators map[string]ValidatorStrategy

func registerDefaultValidators() {
	Register("empty", IsEmpty)
	Register("min", IsMin)
	Register("max", IsMax)
}

func registerValidator(name string, validator ValidatorStrategy) {
	if validators == nil {
		validators = make(map[string]ValidatorStrategy)
	}
	validators[name] = validator
}

func getValidator(name string) (ValidatorStrategy, error) {
	if validators == nil {
		registerDefaultValidators()
	}

	validator, ok := validators[name]

	if !ok {
		return nil, errors.New("Validator '" + name + "' is not registered.")
	}

	return validator, nil
}
