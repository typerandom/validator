package validators

import (
	"github.com/typerandom/validator/core"
	"reflect"
	"time"
)

func EmptyValidator(context core.ValidatorContext, args []interface{}) error {
	if len(args) > 0 {
		return context.NewError("arguments.noneSupported")
	}

	if context.IsNil() {
		return nil
	}

	// TODO: Look into Type.IsZero() and see how much of these "zero" checks it covers.

	switch typedValue := context.Value().(type) {
	case string:
		if len(typedValue) == 0 {
			return nil
		}
	case int64:
		if typedValue == 0 {
			return nil
		}
	case float64:
		if typedValue == 0 {
			return nil
		}
	case bool:
		if typedValue == false {
			return nil
		}
	case time.Time:
		if typedValue.IsZero() {
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
