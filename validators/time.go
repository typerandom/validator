package validators

import (
	"github.com/typerandom/validator/core"
	"time"
)

func TimeValidator(context core.ValidatorContext, options []string) error {
	switch typedValue := context.Value().(type) {
	case string:
		if len(options) != 1 {
			return context.NewError("arguments.singleRequired")
		}

		layout := options[0]

		value, err := time.Parse(layout, typedValue)

		if err != nil {
			return context.NewError("time.mustBeValid")
		}

		if err := context.SetValue(value); err != nil {
			return err
		}

		return nil
	case time.Time:
		if len(options) != 0 {
			return context.NewError("arguments.noneSupported")
		}
		return nil
	}

	return context.NewError("type.unsupported")
}
