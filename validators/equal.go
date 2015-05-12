package validators

import (
	"github.com/typerandom/validator/core"
	"strconv"
)

func EqualValidator(context core.ValidatorContext, args []interface{}) error {
	if len(args) != 1 {
		return context.NewError("arguments.singleRequired")
	}

	if testValue, ok := args[0].(string); ok {
		switch typedValue := context.Value().(type) {
		case string:
			if !context.IsNil() && typedValue == testValue {
				return nil
			}
			return context.NewError("equal.mustEqualValue", testValue)
		case int64:
			parsedTestValue, err := strconv.ParseInt(testValue, 10, 64)

			if err == nil && !context.IsNil() && typedValue == parsedTestValue {
				return nil
			}

			return context.NewError("equal.mustEqualValue", testValue)
		case float64:
			parsedTestValue, err := strconv.ParseFloat(testValue, 64)

			if err == nil && !context.IsNil() && typedValue == parsedTestValue {
				return nil
			}

			return context.NewError("equal.mustEqualValue", testValue)
		case bool:
			parsedTestValue, err := strconv.ParseBool(testValue)

			if err == nil && !context.IsNil() && typedValue == parsedTestValue {
				return nil
			}

			return context.NewError("equal.mustEqualValue", testValue)
		}
	} else {
		return context.NewError("arguments.invalidType", 1, "string")
	}

	return context.NewError("type.unsupported")
}
