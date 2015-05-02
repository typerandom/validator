package validators

import (
	"github.com/typerandom/validator/core"
	"reflect"
	"strconv"
)

func MaxValidator(context core.ValidatorContext, options []string) error {
	if len(options) != 1 {
		return context.NewError("arguments.singleRequired")
	}

	maxValue, err := strconv.ParseFloat(options[0], 64)

	if err != nil {
		return context.NewError("arguments.invalid")
	}

	switch typedValue := context.Value().(type) {
	case string:
		if !context.IsNil() && len(typedValue) > int(maxValue) {
			return context.NewError("max.cannotBeLongerThan", maxValue)
		}
		return nil
	case int64:
		if !context.IsNil() && typedValue > int64(maxValue) {
			return context.NewError("max.cannotBeGreaterThan", maxValue)
		}
		return nil
	case float64:
		if !context.IsNil() && typedValue > maxValue {
			return context.NewError("max.cannotBeGreaterThan", maxValue)
		}
		return nil
	}

	switch context.OriginalKind() {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value()).Len() > int(maxValue) {
			return context.NewError("max.cannotContainMoreItemsThan", maxValue)
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value()).MapKeys()) > int(maxValue) {
			return context.NewError("max.cannotContainMoreKeysThan", maxValue)
		}
		return nil
	}

	return context.NewError("type.unsupported")
}
