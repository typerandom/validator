package core

import (
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
