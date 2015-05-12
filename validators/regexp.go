package validators

import (
	"errors"
	"github.com/typerandom/validator/core"
	"regexp"
)

func RegexpValidator(context core.ValidatorContext, args []interface{}) error {
	if len(args) != 1 {
		return context.NewError("arguments.singleRequired")
	}

	if pattern, ok := args[0].(string); ok {
		if testValue, ok := context.Value().(string); ok {
			if context.IsNil() {
				return context.NewError("regexp.mustMatchPattern", pattern)
			}

			matched, err := regexp.MatchString(pattern, testValue)

			if err != nil {
				return errors.New("Unexpected regexp error for validator field '{field}': " + err.Error())
			}

			if !matched {
				return context.NewError("regexp.mustMatchPattern", pattern)
			}

			return nil
		}
	} else {
		return context.NewError("arguments.invalidType", 1, "string")
	}

	return context.NewError("type.unsupported")
}
