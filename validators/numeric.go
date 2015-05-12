package validators

import (
	"github.com/typerandom/validator/core"
	"strconv"
)

func NumericValidator(context core.ValidatorContext, args []interface{}) error {
	if len(args) > 0 {
		return context.NewError("arguments.noneSupported")
	}

	switch typedValue := context.Value().(type) {
	case string:
		if context.IsNil() || len(typedValue) == 0 {
			return context.NewError("numeric.mustBeNumeric")
		}

		value, err := strconv.ParseFloat(typedValue, 64)

		if err != nil {
			return context.NewError("numeric.mustBeNumeric")
		}

		if err := context.SetValue(value); err != nil {
			return err
		}

		return nil
	case int64:
		return nil
	case float64:
		return nil
	}

	return context.NewError("type.unsupported")
}
