package main

import (
	"errors"
	"reflect"
)

type NormalizedValue struct {
	Value interface{}
	IsNil bool
}

func NewNormalizedValue(value interface{}, isNil bool) *NormalizedValue {
	return &NormalizedValue{
		Value: value,
		IsNil: isNil,
	}
}

// Normalizes all numeric types to their int64 counterparts
func normalizeNumeric(value interface{}) (result interface{}, err error) {
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
		result = float64(typedValue)
	case float64:
		result = typedValue
	default:
		err = errors.New("Unable to resolve value to integer type.")
	}
	return
}

func normalizeValue(value interface{}) (*NormalizedValue, error) {
	isNil := false

	valueType := reflect.ValueOf(value)

	switch valueType.Kind() {
	// Dereference the pointer and normalize that value
	case reflect.Ptr:
		// If it's a nil pointer then flag the value as nil, obtain the inner element and create/return a zero value for the type
		if valueType.IsNil() {
			isNil = true

			value = reflect.Zero(reflect.TypeOf(value).Elem()).Interface()
			normalizedValue, err := normalizeValue(value)

			if err != nil {
				return nil, err
			}

			value = normalizedValue.Value
			break
		}

		value = valueType.Elem().Interface()

		return normalizeValue(value)
	// Normalize all numeric types to their 64-bit counterparts (i.e. int8 -> int64, float32 -> float64)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32:
		var err error
		var normalizedValue interface{}

		if normalizedValue, err = normalizeNumeric(value); err != nil {
			return nil, err
		}

		value = normalizedValue
	}

	return NewNormalizedValue(value, isNil), nil
}
