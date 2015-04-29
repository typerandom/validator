package validator

import (
	"github.com/typerandom/validator/core"
)

var globalInitialized bool
var globalLocale *core.Locale

func assertGlobalInit() {
	if !globalInitialized {
		globalInitialized = true
		if globalLocale == nil {
			globalLocale = core.NewLocale()
			registerDefaultLocale(globalLocale)
			registerDefaultValidators()
		}
	}
}

func registerDefaultLocale(lc *core.Locale) {
	lc.Set("arguments.invalid", "Unable to parse '{validator}' validator options for field '{field}'.")
	lc.Set("arguments.noneSupported", "Validator '{validator}' on field '{field}' does not support any arguments.")
	lc.Set("arguments.singleRequired", "Validator '{validator}' on field '{field}' requires a single argument.")
	lc.Set("arguments.oneOrMoreRequired", "Validator '{validator}' on field '{field}' requires at least one argument.")
	lc.Set("notEmpty.cannotBeEmpty", "{field} cannot be empty.")
	lc.Set("min.cannotBeShorterThan", "{field} cannot be shorter than %d characters.")
	lc.Set("min.cannotBeLessThan", "{field} cannot be less than %d.")
	lc.Set("min.cannotContainLessItemsThan", "{field} cannot contain less than %d items.")
	lc.Set("min.cannotContainLessKeysThan", "{field} cannot contain less than %d keys.")
	lc.Set("max.cannotBeLongerThan", "{field} cannot be longer than %d characters.")
	lc.Set("max.cannotBeGreaterThan", "{field} cannot be greater than %d.")
	lc.Set("max.cannotContainMoreItemsThan", "{field} cannot contain more than %d items.")
	lc.Set("max.cannotContainMoreKeysThan", "{field} cannot contain more than %d keys.")
	lc.Set("lowerCase.mustBeLowerCase", "{field} must be in lower case.")
	lc.Set("upperCase.mustBeUpperCase", "{field} must be in upper case.")
	lc.Set("contain.mustContainValues", "{field} must contain one of the following values '%s'.")
	lc.Set("equal.mustEqualValues", "{field} must equal one of the following values '%s'.")
	lc.Set("regexp.mustMatchPattern", "{field} must match pattern '%s'.")
	lc.Set("numeric.mustBeNumeric", "{field} must be numeric.")
	lc.Set("time.mustBeValid", "{field} must be a valid time.")
}

func registerDefaultValidators() {
	registerValidator("empty", emptyValidator)
	registerValidator("not_empty", notEmptyValidator)
	registerValidator("min", minValidator)
	registerValidator("max", maxValidator)
	registerValidator("lowercase", lowerCaseValidator)
	registerValidator("uppercase", upperCaseValidator)
	registerValidator("contain", containValidator)
	registerValidator("equal", equalValidator)
	registerValidator("regexp", regexpValidator)
	registerValidator("numeric", numericValidator)
	registerValidator("time", timeValidator)
	registerValidator("func", funcValidator)
}

func LoadLocale(jsonPath string) error {
	assertGlobalInit()
	return globalLocale.LoadJson(jsonPath)
}

func Register(name string, validator validatorFilter) {
	registerValidator(name, validator)
}

func Validate(value interface{}) *core.Errors {
	assertGlobalInit()

	context := &Context{
		errors: core.NewErrors(),
		locale: globalLocale,
	}

	walkValidate(context, value, nil)

	return context.errors
}
