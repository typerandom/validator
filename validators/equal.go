package validators

import (
	"github.com/typerandom/validator/core"
	"strings"
)

func EqualValidator(context core.ValidatorContext, options []string) error {
	if len(options) == 0 {
		return context.NewError("arguments.oneOrMoreRequired")
	}

	switch typedValue := context.Value().(type) {
	case string:
		for _, testValue := range options {
			if typedValue == testValue {
				return nil
			}
		}
		return context.NewError("equal.mustEqualValues", strings.Join(options, "', '"))
	}

	return context.NewError("type.unsupported")
}
