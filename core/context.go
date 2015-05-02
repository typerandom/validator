package core

import (
	"errors"
	"reflect"
)

type ValidatorContext interface {
	IsNil() bool
	Value() interface{}
	SetValue(value interface{}) error
	Field() *ReflectedField
	OriginalKind() reflect.Kind
	NewError(localeKey string, args ...interface{}) error
}

type testContext struct {
	value        interface{}
	originalKind reflect.Kind
	isNil        bool
}

func NewTestContext(value interface{}) *testContext {
	ctx := &testContext{}

	if err := ctx.SetValue(value); err != nil {
		return nil
	}

	return ctx
}

func (ctx *testContext) IsNil() bool {
	return ctx.isNil
}

func (ctx *testContext) Value() interface{} {
	return ctx.value
}

func (ctx *testContext) SetValue(value interface{}) error {
	normalized, err := Normalize(value)

	if err != nil {
		return err
	}

	ctx.value = normalized.Value
	ctx.originalKind = normalized.OriginalKind
	ctx.isNil = normalized.IsNil

	return nil
}

func (ctx *testContext) Field() *ReflectedField {
	return nil
}

func (ctx *testContext) OriginalKind() reflect.Kind {
	return ctx.originalKind
}

func (ctx *testContext) NewError(localeKey string, args ...interface{}) error {
	return errors.New(localeKey)
}
