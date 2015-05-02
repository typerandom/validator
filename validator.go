package validator

import (
	"github.com/typerandom/validator/core"
	"github.com/typerandom/validator/validators"
	"sync"
)

type validator struct {
	registry core.ValidatorRegistry
	locale   *core.Locale
	lock     sync.Mutex
}

func (this *validator) LoadLocale(jsonPath string) error {
	return this.locale.LoadJson(jsonPath)
}

func (this *validator) Register(name string, validator core.ValidatorFn) {
	this.registry.Register(name, validator)
}

func (this *validator) Validate(value interface{}) *core.Errors {
	context := &context{
		validator: this,
		errors:    core.NewErrors(),
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

func Validate(value interface{}) *core.Errors {
	return getGlobalValidator().Validate(value)
}
