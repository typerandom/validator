package core

import (
	"errors"
)

type ValidatorFn func(context ValidatorContext, options []string) error

type ValidatorRegistry map[string]ValidatorFn

func NewValidatorRegistry() ValidatorRegistry {
	return make(ValidatorRegistry)
}

func (r ValidatorRegistry) Register(name string, validator ValidatorFn) {
	r[name] = validator
}

func (r ValidatorRegistry) Get(name string) (ValidatorFn, error) {
	validator, ok := r[name]

	if !ok {
		return nil, errors.New("Validator '" + name + "' is not registered.")
	}

	return validator, nil
}
