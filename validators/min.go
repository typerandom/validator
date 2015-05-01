package validators

import (
	"github.com/typerandom/validator/core"
	"reflect"
	"strconv"
)

func MinValidator(context core.ValidatorContext, options []string) error {
	if len(options) != 1 {
		return context.NewError("arguments.singleRequired")
	}

	minValue, err := strconv.Atoi(options[0])

	if err != nil {
		return context.NewError("arguments.invalid")
	}

	switch typedValue := context.Value().(type) {
	case string:
		if context.IsNil() || len(typedValue) < minValue {
			return context.NewError("min.cannotBeShorterThan", minValue)
		}
		return nil
	case int64:
		if context.IsNil() || typedValue < int64(minValue) {
			return context.NewError("min.cannotBeLessThan", minValue)
		}
		return nil
	case float64:
		if context.IsNil() || typedValue < float64(minValue) {
			return context.NewError("min.cannotBeLessThan", minValue)
		}
		return nil
	}

	switch context.OriginalKind() {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value()).Len() < minValue {
			return context.NewError("min.cannotContainLessItemsThan", minValue)
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value()).MapKeys()) < minValue {
			return context.NewError("min.cannotContainLessKeysThan", minValue)
		}
		return nil
	}

	return context.NewError("type.unsupported")
}
