package main

import (
	"errors"
	"reflect"
)

type NormalizedValue struct {
	Value        interface{}
	OriginalKind reflect.Kind
	IsNil        bool
}

func NewNormalizedValue(value interface{}, originalKind reflect.Kind, isNil bool) *NormalizedValue {
	return &NormalizedValue{
		Value:        value,
		OriginalKind: originalKind,
		IsNil:        isNil,
	}
}

// Normalizes all numeric types to their int64 counterparts
func normalizeNumeric(value interface{}) (result interface{}, kind reflect.Kind, err error) {
	kind = reflect.Int64

	switch typedValue := value.(type) {
	case int:
		result = int64(typedValue)
	case int8:
		result = int64(typedValue)
	case int16:
		result = int64(typedValue)
	case int32:
		result = int64(typedValue)
	case int64:
		result = typedValue
	case uint:
		result = int64(typedValue)
	case uint8:
		result = int64(typedValue)
	case uint16:
		result = int64(typedValue)
	case uint32:
		result = int64(typedValue)
	case uint64:
		result = int64(typedValue)
	case float32:
		kind = reflect.Float64
		result = float64(typedValue)
	case float64:
		kind = reflect.Float64
		result = typedValue
	default:
		err = errors.New("Unable to resolve value to integer type.")
	}

	return
}

func normalizeValue(value interface{}, isNil bool) (*NormalizedValue, error) {
	valueType := reflect.ValueOf(value)
	kind := valueType.Kind()

	switch valueType.Kind() {
	// Dereference the pointer and normalize that value
	case reflect.Ptr:
		// If it's a nil pointer then flag the value as nil, obtain the inner element and create/return a zero value for the type
		if valueType.IsNil() {
			isNil = true

			innerElement := reflect.TypeOf(value).Elem()
			kind = innerElement.Kind()

			value = reflect.Zero(innerElement).Interface()
		} else {
			value = valueType.Elem().Interface()
		}

		return normalizeValue(value, isNil)

	// Normalize all numeric types to their 64-bit counterparts (i.e. int8 -> int64, float32 -> float64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32:
		var err error
		var valueKind reflect.Kind
		var normalizedValue interface{}

		if normalizedValue, valueKind, err = normalizeNumeric(value); err != nil {
			return nil, err
		}

		value = normalizedValue
		kind = valueKind
	}

	return NewNormalizedValue(value, kind, isNil), nil
}
