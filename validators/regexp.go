package validators

import (
	"errors"
	"github.com/typerandom/validator/core"
	"regexp"
)

var (
	regexpCache map[string]*regexp.Regexp = map[string]*regexp.Regexp{}
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

			var expr *regexp.Regexp

			if cachedExpr, ok := regexpCache[pattern]; ok {
				expr = cachedExpr
			} else {
				newExpr, err := regexp.Compile(pattern)

				if err != nil {
					return errors.New("Unexpected regexp error for validator field '{field}': " + err.Error())
				}

				expr = newExpr
				regexpCache[pattern] = newExpr
			}

			if !expr.MatchString(testValue) {
				return context.NewError("regexp.mustMatchPattern", pattern)
			}

			return nil
		}
	} else {
		return context.NewError("arguments.invalidType", 1, "string")
	}

	return context.NewError("type.unsupported")
}
