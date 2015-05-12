package validators

import (
	"github.com/typerandom/validator/core"
	"strings"
)

func ContainValidator(context core.ValidatorContext, args []interface{}) error {
	if len(args) != 1 {
		return context.NewError("arguments.singleRequired")
	}

	if testValue, ok := args[0].(string); ok {
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
	} else {
		return context.NewError("arguments.invalidType", 1, "string")
	}

	return context.NewError("type.unsupported")
}
