package main

import (
	"errors"
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
