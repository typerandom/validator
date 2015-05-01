package core

import (
	"errors"
)

type ValidatorFilter func(context ValidatorContext, options []string) error

type ValidatorRegistry map[string]ValidatorFilter

func NewValidatorRegistry() ValidatorRegistry {
	return make(ValidatorRegistry)
}

func (r ValidatorRegistry) Register(name string, validator ValidatorFilter) {
	r[name] = validator
}

func (r ValidatorRegistry) Get(name string) (ValidatorFilter, error) {
	validator, ok := r[name]

	if !ok {
		return nil, errors.New("Validator '" + name + "' is not registered.")
	}

	return validator, nil
}
