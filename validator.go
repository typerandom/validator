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

func (this *validator) Register(name string, validator core.ValidatorFilter) {
	this.registry.Register(name, validator)
}

func (this *validator) Validate(value interface{}) *core.Errors {
	assertGlobalInit()

	context := &context{
		validator: this,
		errors:    core.NewErrors(),
	}

	walkValidate(context, value, nil)

	return context.errors
}

func Default() *validator {
	assertGlobalInit()
	return globalDefaultValidator
}

func New() *validator {
	validator := &validator{
		registry: core.NewValidatorRegistry(),
		locale:   core.NewLocale(),
	}

	validators.RegisterDefaultLocale(validator.locale)
	validators.RegisterDefaultValidators(validator.registry)
	validator.Register("func", funcValidator)

	return validator
}

func Validate(value interface{}) *core.Errors {
	assertGlobalInit()
	return globalDefaultValidator.Validate(value)
}
