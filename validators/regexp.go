package validators

import (
	"errors"
	"github.com/typerandom/validator/core"
	"regexp"
)

func RegexpValidator(context core.ValidatorContext, options []string) error {
	if len(options) != 1 {
		return context.NewError("arguments.singleRequired")
	}

	pattern := options[0]

	if testValue, ok := context.Value().(string); ok {
		matched, err := regexp.MatchString(pattern, testValue)

		if err != nil {
			return errors.New("Unexpected regexp error for validator field '{field}': " + err.Error())
		}

		if !matched {
			return context.NewError("regexp.mustMatchPattern", pattern)
		}

		return nil
	}

	return context.NewError("type.unsupported")
}
