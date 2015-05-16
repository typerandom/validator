package validators

import (
	"fmt"
	"github.com/typerandom/validator/core"
)

func NotValidator(context core.ValidatorContext, args []interface{}) error {
	if len(args) != 1 {
		return context.NewError("arguments.singleRequired")
	}

	var argument = args[0]

	if context.IsNil() {
		if argument == nil {
			return context.NewError("not.cannotBeValue", "'nil'")
		}
		return nil
	}

	switch typedValue := context.Value().(type) {
	case string:
		if typedValue == fmt.Sprintf("%v", argument) {
			return context.NewError("not.cannotBeValue", fmt.Sprintf("'%v'", typedValue))
		}
		return nil
	case int64:
		if typedArgument, ok := argument.(float64); ok {
			if float64(typedValue) == typedArgument {
				return context.NewError("not.cannotBeValue", typedValue)
			}
			return nil
		}
	case float64:
		if typedArgument, ok := argument.(float64); ok {
			if typedValue == typedArgument {
				return context.NewError("not.cannotBeValue", typedValue)
			}
			return nil
		}
	}

	return context.NewError("type.unsupported")
}
