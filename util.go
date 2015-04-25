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

var (
	InvalidMethodError          = errors.New("Method does not exist.")
	InputParameterMismatchError = errors.New("Arguments does not match those of target function.")
	UnhandledCallError          = errors.New("Unhandled function call error.")
)

func CallMethod(i interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)

	// if we start with a pointer, we need to get value pointed to
	// if we start with a value, we need to get a pointer to that value
	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}

	// check for method on value
	method := value.MethodByName(methodName)

	if method.IsValid() {
		finalMethod = method
	}

	// check for method on pointer
	method = ptr.MethodByName(methodName)

	if method.IsValid() {
		finalMethod = method
	} else {
		return nil, InvalidMethodError
	}

	if finalMethod.IsValid() {
		funcType := reflect.TypeOf(finalMethod.Interface())

		numParameters := funcType.NumIn()

		if len(args) != numParameters {
			return nil, InputParameterMismatchError
		}

		for i := 0; i < numParameters; i++ {
			inputType := funcType.In(i)
			if inputType.Kind() != reflect.Interface && reflect.TypeOf(args[i]) != inputType {
				return nil, InputParameterMismatchError
			}
		}

		methodArgs := make([]reflect.Value, len(args))

		for i, arg := range args {
			methodArgs[i] = reflect.ValueOf(arg)
		}

		callResult := finalMethod.Call(methodArgs)

		returnValues := make([]interface{}, len(callResult))

		for i, result := range callResult {
			returnValues[i] = result.Interface()
		}

		return returnValues, nil
	}

	// return or panic, method not found of either type
	return nil, UnhandledCallError
}
