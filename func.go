package validator

import (
	"errors"
	"github.com/typerandom/validator/core"
)

func funcValidator(ctx core.ValidatorContext, options []string) error {
	var funcName string

	switch len(options) {
	case 0:
		funcName = "Validate" + ctx.Field().Name
	case 1:
		funcName = options[0]
	default:
		return ctx.NewError("arguments.singleRequired")
	}

	if internalContext, ok := ctx.(*context); ok {
		returnValues, err := core.CallDynamicMethod(internalContext.source, funcName, ctx)

		if err != nil {
			if err == core.InvalidMethodError {
				return errors.New("Validation method '" + ctx.Field().Parent.FullName(funcName) + "' on field '{field}' does not exist.")
			}
			return err
		}

		if len(returnValues) == 1 {
			if returnValues[0] == nil {
				return nil
			} else if err, ok := returnValues[0].(error); ok {
				return errors.New(err.Error())
			} else {
				return errors.New("Invalid return value of validation method '" + ctx.Field().Parent.FullName(funcName) + "'. Return value must be of type 'error'.")
			}
		}
	} else {
		return errors.New("Invalid context provided.")
	}

	return errors.New("Validator not supported.")
}
