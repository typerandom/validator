package validators

import (
	"github.com/typerandom/validator/core"
	"time"
)

func TimeValidator(context core.ValidatorContext, args []interface{}) error {
	switch typedValue := context.Value().(type) {
	case string:
		if len(args) != 1 {
			return context.NewError("arguments.singleRequired")
		}

		if layout, ok := args[0].(string); ok {
			value, err := time.Parse(layout, typedValue)

			if err != nil {
				return context.NewError("time.mustBeValid")
			}

			if err := context.SetValue(value); err != nil {
				return err
			}

			return nil
		} else {
			return context.NewError("arguments.invalidType", 1, "string")
		}
	case time.Time:
		return nil
	}

	return context.NewError("type.unsupported")
}
