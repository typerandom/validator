package validators

import (
	"github.com/typerandom/validator/core"
	"strings"
)

func ContainValidator(context core.ValidatorContext, options []string) error {
	if len(options) != 1 {
		return context.NewError("arguments.singleRequired")
	}

	testValue := options[0]

	switch typedValue := context.Value().(type) {
	case string:
		if len(testValue) == 0 {
			return context.NewError("arguments.invalid")
		}

		if context.IsNil() || !strings.Contains(typedValue, testValue) {
			return context.NewError("contain.mustContainValue", testValue)
		}

		return nil
	}

	/*TODO: Add support for checking if a value is present in a slice/array/map.*/

	return context.NewError("type.unsupported")
}
