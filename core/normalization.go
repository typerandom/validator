package core

import (
	"reflect"
)

type NormalizedValue struct {
	Value        interface{}
	OriginalKind reflect.Kind
	IsNil        bool
}

// TODO: Normalize slices to arrays?

func normalizeInternal(value interface{}, isNil bool) (*NormalizedValue, error) {
	reflectedValue := reflect.ValueOf(value)
	kind := reflectedValue.Kind()

	switch reflectedValue.Kind() {

	// Dereference the pointer and normalize that value
	case reflect.Ptr:
		// If it's a nil pointer then flag the value as nil, obtain the inner element and create/return a zero value for the type
		if reflectedValue.IsNil() {
			isNil = true

			innerElement := reflect.TypeOf(value).Elem()
			kind = innerElement.Kind()

			value = reflect.Zero(innerElement).Interface()
		} else {
			value = reflectedValue.Elem().Interface()
		}

		return normalizeInternal(value, isNil)

	// Convert any number to its 64-bit counterpart. Also normalize according to kind. I.e. any string, int, float or bool kind will be normalized to its base type.
	// This means that a type, lets say `type Id int64` would instead of having the type `main.Id` be normalized to `int64`.
	// This is done in order to simplify work for the validators so that they don't have to validate by kind or have to care for custom types.
	// It's also easier for them to validate if the expected type is always the same, i.e. int64 instead of int8, uint16...

	case reflect.String:
		value = reflectedValue.String()

	case reflect.Bool:
		value = reflectedValue.Bool()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		value = int64(reflectedValue.Uint())

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = reflectedValue.Int()

	case reflect.Float32, reflect.Float64:
		value = reflectedValue.Float()

	case reflect.Invalid:
		if value == nil {
			isNil = true
		}
	}

	normalized := &NormalizedValue{
		Value:        value,
		OriginalKind: kind,
		IsNil:        isNil,
	}

	return normalized, nil
}

func Normalize(value interface{}) (*NormalizedValue, error) {
	return normalizeInternal(value, false)
}
