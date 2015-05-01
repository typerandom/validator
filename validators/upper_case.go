package validators

import (
	"github.com/typerandom/validator/core"
	"unicode"
)

func UpperCaseValidator(context core.ValidatorContext, options []string) error {
	if len(options) > 0 {
		return context.NewError("arguments.noneSupported")
	}

	switch typedValue := context.Value().(type) {
	case string:
		if context.IsNil() || len(typedValue) == 0 {
			return nil
		}

		for _, char := range typedValue {
			if unicode.IsLetter(char) && !unicode.IsUpper(char) {
				return context.NewError("upperCase.mustBeUpperCase")
			}
		}

		return nil
	}

	return context.NewError("type.unsupported")
}
