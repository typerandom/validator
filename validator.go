// Package validator provides validation methods for validating structures.
package validator

import (
	"github.com/typerandom/validator/core"
	"github.com/typerandom/validator/validators"
	"sync"
)

// Validator represents a validator with it's own configuration set.
type Validator struct {
	// The tag that is used for the field's display name.
	// Default: Empty string that defaults to the field name.
	DisplayNameTag string

	registry core.ValidatorRegistry
	locale   *core.Locale
	lock     sync.Mutex
}

// Locale retrieves the locale for this validator.
func (this *Validator) Locale() *core.Locale {
	return this.locale
}

// Register registers a validator by name.
func (this *Validator) Register(name string, validator core.ValidatorFn) {
	this.registry.Register(name, validator)
}

// Validate validates fields of a structure, or structures of a map, slice or array.
func (this *Validator) Validate(value interface{}) core.ErrorList {
	context := &context{
		validator: this,
	}

	walkValidate(context, value, nil)

	return context.errors
}

// CheckSyntax checks the validate tag syntax of a structure.
func CheckSyntax(value interface{}) error {
	if _, err := core.GetStructFields(value, "validator", ""); err != nil {
		return err
	}
	return nil
}

// New creates a new validator.
func New() *Validator {
	validator := &Validator{
		registry: core.NewValidatorRegistry(),
		locale:   core.NewLocale(),
	}

	validators.RegisterDefaultLocale(validator.locale)
	validators.RegisterDefaultValidators(validator.registry)

	return validator
}

// Default retrieves the default global validator.
// This is what is used when you call the global Validate or Register method.
func Default() *Validator {
	return getGlobalValidator()
}

// Register registers a validator method by name.
func Register(name string, validator core.ValidatorFn) {
	getGlobalValidator().Register(name, validator)
}

// Validate validates fields of a structure, or structures of a map, slice or array.
func Validate(value interface{}) core.ErrorList {
	return getGlobalValidator().Validate(value)
}
