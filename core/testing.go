package core

import (
	"errors"
	"reflect"
)

type testContext struct {
	source interface{}

	value        interface{}
	originalKind reflect.Kind
	isNil        bool

	field *ReflectedField
}

func NewTestContext(value interface{}) *testContext {
	ctx := &testContext{}

	if err := ctx.SetValue(value); err != nil {
		return nil
	}

	return ctx
}

func (this *testContext) SetSource(source interface{}) {
	this.source = source
}

func (this *testContext) Source() interface{} {
	return this.source
}

func (this *testContext) IsNil() bool {
	return this.isNil
}

func (this *testContext) Value() interface{} {
	return this.value
}

func (this *testContext) SetValue(value interface{}) error {
	normalized, err := Normalize(value)

	if err != nil {
		return err
	}

	this.value = normalized.Value
	this.originalKind = normalized.OriginalKind
	this.isNil = normalized.IsNil

	return nil
}

func (this *testContext) SetField(field *ReflectedField) {
	this.field = field
}

func (this *testContext) Field() *ReflectedField {
	return this.field
}

func (this *testContext) OriginalKind() reflect.Kind {
	return this.originalKind
}

func (this *testContext) NewError(localeKey string, args ...interface{}) error {
	return errors.New(localeKey)
}
