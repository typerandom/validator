// Package validator provides validation methods for validating structures.
package validator

import (
	"github.com/typerandom/validator/core"
	"github.com/typerandom/validator/validators"
	"sync"
)

type Validator interface {
	// The tag that is used for the field's display name.
	// Default: Empty string that defaults to the field name.
	SetDisplayNameTag(name string)

	// Locale retrieves the locale for this validator.
	Locale() *core.Locale

	// Register registers a validator by name.
	Register(name string, validator core.ValidatorFn)

	// Validate validates fields of a structure, or structures of a map, slice or array.
	Validate(value interface{}) core.ErrorList
}

// Validator represents a validator with it's own configuration set.
type validator struct {
	displayNameTag *string

	registry core.ValidatorRegistry
	locale   *core.Locale
	lock     sync.Mutex
}

func newValidator() *validator {
	validator := &validator{
		registry: core.NewValidatorRegistry(),
		locale:   core.NewLocale(),
	}

	validators.RegisterDefaultLocale(validator.locale)
	validators.RegisterDefaultValidators(validator.registry)

	return validator
}

func (this *validator) Locale() *core.Locale {
	return this.locale
}

func (this *validator) SetDisplayNameTag(tagName string) {
	this.displayNameTag = &tagName
}

func (this *validator) Register(name string, validator core.ValidatorFn) {
	this.registry.Register(name, validator)
}

func (this *validator) Validate(value interface{}) core.ErrorList {
	context := &context{
		validator: this,
	}

	walkValidate(context, value, nil)

	return context.errors
}

// CheckSyntax checks the validate tag syntax of a structure.
func CheckSyntax(value interface{}) error {
	if _, err := core.GetStructFields(value, "validator", nil); err != nil {
		return err
	}
	return nil
}

// New creates a new validator.
func New() Validator {
	return newValidator()
}

// Default retrieves the default global validator.
// It is a singleton method that always returns the same instance.
// This is the validator that is used when you call the global Validate() or Register() method.
func Default() Validator {
	return getGlobalValidator()
}

// Register registers a validator method by name on the default validator.
func Register(name string, validator core.ValidatorFn) {
	getGlobalValidator().Register(name, validator)
}

// Validate validates fields of a structure, or structures of a map, slice or array using the default validator.
func Validate(value interface{}) core.ErrorList {
	return getGlobalValidator().Validate(value)
}
