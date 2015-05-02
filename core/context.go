package core

import (
	"reflect"
)

type ValidatorContext interface {
	// Source returns the object from which the field is referenced from.
	Source() interface{}

	// Field returns the field from the struct which this value was referenced from.
	Field() *ReflectedField

	// Value returns the normalized value.
	Value() interface{}

	// SetValue normalizes the value and sets value, originalKind and nil of the context.
	// Returns error if the value fails to be normalized.
	SetValue(value interface{}) error

	// IsNil indicates whether or not the value was nil before normalization.
	// I.e. because normalization removes pointers from value.
	IsNil() bool

	// OriginalKind returns the kind (non pointer) before value type was normalized.
	// I.e. if the type of the value set was *int8, then the OriginalKind would be int8.
	OriginalKind() reflect.Kind

	// NewError returns a formatted error based on a locale key and format parameters.
	// If the locale key does not exist, then an error is returned.
	NewError(localeKey string, args ...interface{}) error
}
