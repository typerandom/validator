package cocoon

import (
	"errors"
	"github.com/typerandom/cocoon/core"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func newValidatorError(field *core.ReflectedField, tag *core.Tag, err error) *ValidatorError {
	return &ValidatorError{
		Field:  field,
		Tag:    tag,
		Source: err,
	}
}

type ValidatorError struct {
	Field  *core.ReflectedField
	Tag    *core.Tag
	Source error
}

func (this *ValidatorError) String() string {
	return "{ error: " + this.Error() + "}"
}

func (this *ValidatorError) Error() string {
	message := strings.Replace(this.Source.Error(), "{field}", this.Field.FullName(), 1)
	message = strings.Replace(message, "{validator}", this.Tag.Name, 1)
	return message
}

type validatorFilter func(context *Context, options []string) error

var validators map[string]validatorFilter

func registerValidator(name string, validator validatorFilter) {
	if validators == nil {
		validators = make(map[string]validatorFilter)
	}
	validators[name] = validator
}

func getValidator(name string) (validatorFilter, error) {
	validator, ok := validators[name]

	if !ok {
		return nil, errors.New("Validator '" + name + "' is not registered.")
	}

	return validator, nil
}

var UnsupportedTypeError = errors.New("Validator not supported.")

func emptyValidator(context *Context, options []string) error {
	if len(options) > 0 {
		return context.GetLocalizedError("arguments.noneSupported")
	}

	if context.IsNil {
		context.StopValidate = true
		return nil
	}

	switch typedValue := context.Value.(type) {
	case string:
		if len(typedValue) == 0 {
			context.StopValidate = true
		}
	case int64:
		if typedValue == 0 {
			context.StopValidate = true
		}
	}

	switch context.OriginalKind {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value).Len() == 0 {
			context.StopValidate = true
		}
	case reflect.Map:
		if len(reflect.ValueOf(context.Value).MapKeys()) == 0 {
			context.StopValidate = true
		}
	}

	return nil
}

func notEmptyValidator(context *Context, options []string) error {
	if len(options) > 0 {
		return context.GetLocalizedError("arguments.noneSupported")
	}

	cannotBeEmptyError := func() error {
		return context.GetLocalizedError("notEmpty.cannotBeEmpty")
	}

	if context.IsNil {
		return cannotBeEmptyError()
	}

	switch typedValue := context.Value.(type) {
	case string:
		if len(typedValue) == 0 {
			return cannotBeEmptyError()
		}
	case int64:
		if typedValue == 0 {
			return cannotBeEmptyError()
		}
	case float64:
		if typedValue == 0 {
			return cannotBeEmptyError()
		}
	}

	switch context.OriginalKind {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value).Len() == 0 {
			return cannotBeEmptyError()
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value).MapKeys()) == 0 {
			return cannotBeEmptyError()
		}
	}

	return nil
}

func minValidator(context *Context, options []string) error {
	if len(options) != 1 {
		return context.GetLocalizedError("arguments.singleRequired")
	}

	minValue, err := strconv.Atoi(options[0])

	if err != nil {
		return context.GetLocalizedError("arguments.invalid")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if context.IsNil || len(typedValue) < minValue {
			return context.GetLocalizedError("min.cannotBeShorterThan", minValue)
		}
		return nil
	case int64:
		if context.IsNil || typedValue < int64(minValue) {
			return context.GetLocalizedError("min.cannotBeLessThan", minValue)
		}
		return nil
	case float64:
		if context.IsNil || typedValue < float64(minValue) {
			return context.GetLocalizedError("min.cannotBeLessThan", minValue)
		}
		return nil
	}

	switch context.OriginalKind {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value).Len() < minValue {
			return context.GetLocalizedError("min.cannotContainLessItemsThan", minValue)
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value).MapKeys()) < minValue {
			return context.GetLocalizedError("min.cannotContainLessKeysThan", minValue)
		}
		return nil
	}

	return UnsupportedTypeError
}

func maxValidator(context *Context, options []string) error {
	if len(options) != 1 {
		return context.GetLocalizedError("arguments.singleRequired")
	}

	maxValue, err := strconv.Atoi(options[0])

	if err != nil {
		return context.GetLocalizedError("arguments.invalid")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if !context.IsNil && len(typedValue) > maxValue {
			return context.GetLocalizedError("max.cannotBeLongerThan", maxValue)
		}
		return nil
	case int64:
		if !context.IsNil && typedValue > int64(maxValue) {
			return context.GetLocalizedError("max.cannotBeGreaterThan", maxValue)
		}
		return nil
	case float64:
		if !context.IsNil && typedValue > float64(maxValue) {
			return context.GetLocalizedError("max.cannotBeGreaterThan", maxValue)
		}
		return nil
	}

	switch context.OriginalKind {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value).Len() > maxValue {
			return context.GetLocalizedError("max.cannotContainMoreItemsThan", maxValue)
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value).MapKeys()) > maxValue {
			return context.GetLocalizedError("max.cannotContainMoreKeysThan", maxValue)
		}
		return nil
	}

	return UnsupportedTypeError
}

func equalValidator(context *Context, options []string) error {
	if len(options) == 0 {
		return context.GetLocalizedError("arguments.oneOrMoreRequired")
	}

	switch typedValue := context.Value.(type) {
	case string:
		for _, testValue := range options {
			if typedValue == testValue {
				return nil
			}
		}
		return context.GetLocalizedError("equal.mustEqualValues", strings.Join(options, "', '"))
	}

	return UnsupportedTypeError
}

func containValidator(context *Context, options []string) error {
	if len(options) == 0 {
		return context.GetLocalizedError("arguments.oneOrMoreRequired")
	}

	switch typedValue := context.Value.(type) {
	case string:
		for _, testValue := range options {
			if !strings.Contains(typedValue, testValue) {
				return context.GetLocalizedError("contain.mustContainValues", strings.Join(options, "', '"))
			}
		}
		return nil
	}

	return UnsupportedTypeError
}

func lowerCaseValidator(context *Context, options []string) error {
	if len(options) > 0 {
		return context.GetLocalizedError("arguments.noneSupported")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if context.IsNil || len(typedValue) == 0 {
			return nil
		}

		for _, char := range typedValue {
			if unicode.IsLetter(char) && !unicode.IsLower(char) {
				return context.GetLocalizedError("lowerCase.mustBeLowerCase")
			}
		}

		return nil
	}

	return UnsupportedTypeError
}

func upperCaseValidator(context *Context, options []string) error {
	if len(options) > 0 {
		return context.GetLocalizedError("arguments.noneSupported")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if context.IsNil || len(typedValue) == 0 {
			return nil
		}

		for _, char := range typedValue {
			if unicode.IsLetter(char) && !unicode.IsUpper(char) {
				return context.GetLocalizedError("upperCase.mustBeUpperCase")
			}
		}

		return nil
	}

	return UnsupportedTypeError
}

func numericValidator(context *Context, options []string) error {
	if len(options) > 0 {
		return context.GetLocalizedError("arguments.noneSupported")
	}

	switch typedValue := context.Value.(type) {
	case string:
		if context.IsNil || len(typedValue) == 0 {
			return context.GetLocalizedError("numeric.mustBeNumeric")
		}

		value, err := strconv.ParseInt(typedValue, 10, 32)

		if err != nil {
			return context.GetLocalizedError("numeric.mustBeNumeric")
		}

		context.Value = value

		return nil
	case int64:
		return nil
	case float64:
		return nil
	}

	return UnsupportedTypeError
}

func timeValidator(context *Context, options []string) error {
	switch typedValue := context.Value.(type) {
	case string:
		if len(options) != 1 {
			return context.GetLocalizedError("arguments.singleRequired")
		}

		layout := options[0]

		value, err := time.Parse(layout, typedValue)

		if err != nil {
			return context.GetLocalizedError("time.mustBeValid")
		}

		context.Value = value

		return nil
	case time.Time:
		if len(options) != 0 {
			return context.GetLocalizedError("arguments.noneSupported")
		}
		return nil
	}

	return UnsupportedTypeError
}

func regexpValidator(context *Context, options []string) error {
	if len(options) != 1 {
		return context.GetLocalizedError("arguments.singleRequired")
	}

	pattern := options[0]

	if testValue, ok := context.Value.(string); ok {
		matched, err := regexp.MatchString(pattern, testValue)

		if err != nil {
			return errors.New("Unexpected regexp error for validator field '{field}': " + err.Error())
		}

		if !matched {
			return context.GetLocalizedError("regexp.mustMatchPattern", pattern)
		}

		return nil
	}

	return UnsupportedTypeError
}

func funcValidator(context *Context, options []string) error {
	var funcName string

	switch len(options) {
	case 0:
		funcName = "Validate" + context.Field.Name
	case 1:
		funcName = options[0]
	default:
		return context.GetLocalizedError("arguments.singleRequired")
	}

	returnValues, err := core.CallDynamicMethod(context.source, funcName, context)

	if err != nil {
		if err == core.InvalidMethodError {
			return errors.New("Validation method '" + context.Field.Parent.FullName(funcName) + "' on field '{field}' does not exist.")
		}
		return err
	}

	if len(returnValues) == 1 {
		if returnValues[0] == nil {
			return nil
		} else if err, ok := returnValues[0].(error); ok {
			return errors.New(err.Error())
		} else {
			return errors.New("Invalid return value of validation method '" + context.Field.Parent.FullName(funcName) + "'. Return value must be of type 'error'.")
		}
	}

	return errors.New("Validator not supported.")
}
