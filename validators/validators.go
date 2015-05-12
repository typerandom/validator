package validators

import (
	"github.com/typerandom/validator/core"
)

func RegisterDefaultLocale(lc *core.Locale) {
	lc.Set("type.unsupported", "Validator '{validator}' does not support the type of field '{field}'.")
	lc.Set("arguments.invalid", "Unable to parse '{validator}' validator options for field '{field}'.")
	lc.Set("arguments.invalidType", "Validator '{validator}' on field '{field}' requires parameter %d to be of type %s.")
	lc.Set("arguments.noneSupported", "Validator '{validator}' on field '{field}' does not support any arguments.")
	lc.Set("arguments.singleRequired", "Validator '{validator}' on field '{field}' requires a single argument.")
	lc.Set("arguments.oneOrMoreRequired", "Validator '{validator}' on field '{field}' requires at least one argument.")
	lc.Set("nil.isNotNil", "{field} is not nil.")
	lc.Set("empty.isNotEmpty", "{field} is not empty.")
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
	lc.Set("contain.mustContainValue", "{field} must contain one of the following values '%s'.")
	lc.Set("equal.mustEqualValue", "{field} must equal one of the following values '%s'.")
	lc.Set("regexp.mustMatchPattern", "{field} must match pattern '%s'.")
	lc.Set("numeric.mustBeNumeric", "{field} must be numeric.")
	lc.Set("time.mustBeValid", "{field} must be a valid time.")
}

func RegisterDefaultValidators(r core.ValidatorRegistry) {
	r.Register("nil", NilValidator)
	r.Register("empty", EmptyValidator)
	r.Register("not_empty", NotEmptyValidator)
	r.Register("min", MinValidator)
	r.Register("max", MaxValidator)
	r.Register("lowercase", LowerCaseValidator)
	r.Register("uppercase", UpperCaseValidator)
	r.Register("contain", ContainValidator)
	r.Register("equal", EqualValidator)
	r.Register("regexp", RegexpValidator)
	r.Register("numeric", NumericValidator)
	r.Register("time", TimeValidator)
	r.Register("func", FuncValidator)
}
