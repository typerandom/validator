package validator

import (
	"github.com/typerandom/validator/core"
	"github.com/typerandom/validator/validators"
	"sync"
)

type validator struct {
	// The tag that is used for the field's display name.
	// Default: Empty string that defaults to the field name.
	DisplayNameTag string

	registry core.ValidatorRegistry
	locale   *core.Locale
	lock     sync.Mutex
}

func (this *validator) Locale() *core.Locale {
	return this.locale
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

func New() *validator {
	validator := &validator{
		registry: core.NewValidatorRegistry(),
		locale:   core.NewLocale(),
	}

	validators.RegisterDefaultLocale(validator.locale)
	validators.RegisterDefaultValidators(validator.registry)

	return validator
}

func Default() *validator {
	return getGlobalValidator()
}

func Register(name string, validator core.ValidatorFn) {
	getGlobalValidator().Register(name, validator)
}

func Validate(value interface{}) core.ErrorList {
	return getGlobalValidator().Validate(value)
}
