package main

import (
	"errors"
	"reflect"
)

// Global validator registry

var validators map[string]ValidatorFn

func registerValidator(name string, validator ValidatorFn) {
	if validators == nil {
		validators = make(map[string]ValidatorFn)
	}
	validators[name] = validator
}

func getValidator(name string) (ValidatorFn, error) {
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

var structFieldValidatorRegistry map[reflect.Type]map[string][]ValidatorFn

func registerStructFieldValidators(structType reflect.Type, validators map[string][]ValidatorFn) {
	if structFieldValidatorRegistry == nil {
		structFieldValidatorRegistry = make(map[reflect.Type]map[string][]ValidatorFn)
	}

	structFieldValidatorRegistry[structType] = validators
}
