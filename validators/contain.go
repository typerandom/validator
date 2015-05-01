package validators

import (
	"github.com/typerandom/validator/core"
	"strings"
)

func ContainValidator(context core.ValidatorContext, options []string) error {
	if len(options) == 0 {
		return context.NewError("arguments.oneOrMoreRequired")
	}

	switch typedValue := context.Value().(type) {
	case string:
		for _, testValue := range options {
			if !strings.Contains(typedValue, testValue) {
				return context.NewError("contain.mustContainValues", strings.Join(options, "', '"))
			}
		}
		return nil
	}

	return context.NewError("type.unsupported")
}
