package validators

import (
	"errors"
	"github.com/typerandom/validator/core"
)

func FuncValidator(context core.ValidatorContext, args []interface{}) error {
	var funcName string
	var funcArgs []interface{}

	if len(args) == 0 {
		funcName = "Validate" + context.Field().Name
	} else {
		if typedArg, ok := args[0].(string); ok {
			funcName = typedArg
			if len(args) > 1 {
				funcArgs = args[1:]
			}
		} else {
			return context.NewError("arguments.invalidType", 1, "string")
		}
	}

	returnValues, err := core.CallDynamicMethod(context.Source(), funcName, context, funcArgs)

	if err != nil {
		if err == core.InvalidMethodError {
			return errors.New("Validation method '" + context.Field().Parent.FullName(funcName) + "' on field '{field}' does not exist.")
		}
		return err
	}

	if len(returnValues) == 1 {
		if returnValues[0] == nil {
			return nil
		} else if err, ok := returnValues[0].(error); ok {
			return err
		}
	}

	return errors.New("Invalid return value(s) of validation method '" + context.Field().Parent.FullName(funcName) + "'. Return value must be of type 'error'.")
}
