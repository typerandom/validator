package validators

import (
	"github.com/typerandom/validator/core"
)

func NilValidator(context core.ValidatorContext, options []string) error {
	if len(options) > 0 {
		return context.NewError("arguments.noneSupported")
	}

	if context.IsNil() {
		return nil
	}

	return context.NewError("nil.isNotNil")
}
