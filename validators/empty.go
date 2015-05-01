package validators

import (
	"github.com/typerandom/validator/core"
	"reflect"
)

func EmptyValidator(context core.ValidatorContext, options []string) error {
	if len(options) > 0 {
		return context.NewError("arguments.noneSupported")
	}

	if context.IsNil() {
		return nil
	}

	switch typedValue := context.Value().(type) {
	case string:
		if len(typedValue) == 0 {
			return nil
		}
	case int64:
		if typedValue == 0 {
			return nil
		}
	}

	switch context.OriginalKind() {
	case reflect.Array, reflect.Slice:
		if reflect.ValueOf(context.Value()).Len() == 0 {
			return nil
		}
	case reflect.Map:
		if len(reflect.ValueOf(context.Value()).MapKeys()) == 0 {
			return nil
		}
	}

	return context.NewError("empty.isNotEmpty")
}
