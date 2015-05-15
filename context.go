package validator

import (
	"errors"
	"fmt"
	"github.com/typerandom/validator/core"
	"reflect"
)

type context struct {
	validator *Validator

	value        interface{}
	originalKind reflect.Kind
	field        *core.ReflectedField
	isNil        bool

	errors core.ErrorList
	source interface{}
}

func (this *context) Source() interface{} {
	return this.source
}

func (this *context) Value() interface{} {
	return this.value
}

func (this *context) SetValue(value interface{}) error {
	normalized, err := core.Normalize(value)

	if err != nil {
		return err
	}

	this.value = normalized.Value
	this.originalKind = normalized.OriginalKind
	this.isNil = normalized.IsNil

	return nil
}

func (this *context) OriginalKind() reflect.Kind {
	return this.originalKind
}

func (this *context) Field() *core.ReflectedField {
	return this.field
}

func (this *context) IsNil() bool {
	return this.isNil
}

func (this *context) NewError(localeKey string, args ...interface{}) error {
	message, err := this.validator.locale.Get(localeKey)

	if err != nil {
		return err
	}

	if len(args) > 0 {
		message = fmt.Sprintf(message, args...)
	}

	return errors.New(message)
}

func (this *context) setValue(normalized *core.NormalizedValue) {
	this.value = normalized.Value
	this.originalKind = normalized.OriginalKind
	this.isNil = normalized.IsNil
}

func (this *context) setSource(source interface{}) {
	this.source = source
}

func (this *context) setField(field *core.ReflectedField) {
	this.field = field
}
