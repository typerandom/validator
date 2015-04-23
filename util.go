package main

import (
	"errors"
	"reflect"
)

// If integer, then tries to resolve the value to either a int64 value or pointer.
// If it fails to do that, then the original value is returned.
// Note: This could have been done with reflection. But this is probably clearer and better performance wise.
// ALSO: Note that floats will be truncated according to http://golang.org/ref/spec#Conversions (float -> int)
func normalizeToInt64(value interface{}) (result interface{}, err error) {
	switch typedValue := value.(type) {
	case int:
		result = int64(typedValue)
	case int8:
		result = int64(typedValue)
	case int16:
		result = int64(typedValue)
	case int32:
		result = int64(typedValue)
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
	// Note that converting a float to int will discard any fraction.
	// It's important that this effect is documented.
	case float32:
		result = int64(typedValue)
	case float64:
		result = int64(typedValue)
	default:
		err = errors.New("Unable to resolve value to integer type.")
	}

	return
}

// Check into http://stackoverflow.com/questions/18091562/how-to-get-underlying-value-from-a-reflect-value-in-golang

func normalizeValue(value interface{}) (interface{}, bool, error) {
	valueType := reflect.ValueOf(value)

	switch valueType.Kind() {
	// Dereference the pointer and normalize that value
	case reflect.Ptr:
		if valueType.IsNil() {
			return value, true, nil
		}

		value = valueType.Elem().Interface()

		return normalizeValue(value)
	// Reflect all numeric types except reflect.Int64, since that is what we want to resolve to
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64:
		var err error
		var normalizedValue interface{}

		if normalizedValue, err = normalizeToInt64(value); err != nil {
			return nil, false, err
		}

		value = normalizedValue
	}

	return value, false, nil
}
