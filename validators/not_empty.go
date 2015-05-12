package validators

import (
	"github.com/typerandom/validator/core"
	"reflect"
)

func NotEmptyValidator(context core.ValidatorContext, args []interface{}) error {
	if len(args) > 0 {
		return context.NewError("arguments.noneSupported")
	}

	cannotBeEmptyError := func() error {
		return context.NewError("notEmpty.cannotBeEmpty")
	}

	if context.IsNil() {
		return cannotBeEmptyError()
	}

	switch typedValue := context.Value().(type) {
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

	switch context.OriginalKind() {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value()).Len() == 0 {
			return cannotBeEmptyError()
		}
		return nil
	case reflect.Map:
		if len(reflect.ValueOf(context.Value()).MapKeys()) == 0 {
			return cannotBeEmptyError()
		}
	}

	return nil
}
