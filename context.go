package cocoon

import (
	"errors"
	"fmt"
	"github.com/typerandom/cocoon/core"
	"reflect"
)

type Context struct {
	Value        interface{}
	OriginalKind reflect.Kind
	Field        *core.ReflectedField
	IsNil        bool
	StopValidate bool

	locale *core.Locale
	errors *core.Errors
	source interface{}
}

func (this *Context) setValue(normalized *core.NormalizedValue) {
	this.Value = normalized.Value
	this.OriginalKind = normalized.OriginalKind
	this.IsNil = normalized.IsNil
}

func (this *Context) setSource(source interface{}) {
	this.source = source
}

func (this *Context) setField(field *core.ReflectedField) {
	this.Field = field
}

func (this *Context) GetLocalizedError(key string, args ...interface{}) error {
	message, err := this.locale.Get(key)

	if err != nil {
		return err
	}

	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	return errors.New(message)
}
