package validators

import (
	"errors"
	"github.com/typerandom/validator/core"
)

func FuncValidator(context core.ValidatorContext, options []string) error {
	var funcName string

	switch len(options) {
	case 0:
		funcName = "Validate" + context.Field().Name
	case 1:
		funcName = options[0]
	default:
		return context.NewError("arguments.singleRequired")
	}

	returnValues, err := core.CallDynamicMethod(context.Source(), funcName, context)

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
