package main

import (
	"errors"
	"reflect"
)

// Global validator registry

var validators map[string]ValidatorFilter

func registerValidator(name string, validator ValidatorFilter) {
	if validators == nil {
		validators = make(map[string]ValidatorFilter)
	}
	validators[name] = validator
}

func getValidator(name string) (ValidatorFilter, error) {
	if validators == nil {
		registerDefaultValidators()
	}

	validator, ok := validators[name]

	if !ok {
		return nil, errors.New("Validator '" + name + "' is not registered.")
	}

	return validator, nil
}

// Type validator registry

var structFieldValidatorRegistry map[reflect.Type]map[string][]ValidatorFilter

func registerStructFieldValidators(structType reflect.Type, validators map[string][]ValidatorFilter) {
	if structFieldValidatorRegistry == nil {
		structFieldValidatorRegistry = make(map[reflect.Type]map[string][]ValidatorFilter)
	}

	structFieldValidatorRegistry[structType] = validators
}
