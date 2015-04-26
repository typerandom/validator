package main

import (
	"errors"
	"reflect"
)

var (
	InvalidMethodError          = errors.New("Method does not exist.")
	InputParameterMismatchError = errors.New("Arguments does not match those of target function.")
	UnhandledCallError          = errors.New("Unhandled function call error.")
)

func callMethod(i interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
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
