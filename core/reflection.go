package core

import (
	"errors"
	"github.com/typerandom/validator/core/parser"
	"reflect"
	"strings"
	"unicode"
)

type ReflectedField struct {
	Index        int
	Parent       *ReflectedField
	Name         string
	MethodGroups []parser.Methods
}

func (this *ReflectedField) GetValue(sourceStruct reflect.Value) interface{} {
	return sourceStruct.Field(this.Index).Interface()
}

func (this *ReflectedField) FullName(postfix ...string) string {
	fullName := this.Name
	parent := this.Parent

	for parent != nil {
		if len(parent.Name) > 0 {
			fullName = parent.Name + "." + fullName
		}
		parent = parent.Parent
	}

	if len(postfix) > 0 {
		if len(fullName) > 0 {
			fullName += "."
		}
		fullName += strings.Join(postfix, ".")
	}

	return fullName
}

func reflectValue(value interface{}) reflect.Type {
	reflectedValueType := reflect.TypeOf(value)

	if reflectedValueType.Kind() == reflect.Ptr {
		reflectedValueType = reflectedValueType.Elem()
	}

	return reflectedValueType
}

var structFieldCache map[reflect.Type][]*ReflectedField = map[reflect.Type][]*ReflectedField{}

func GetStructFields(value interface{}, tagName string) ([]*ReflectedField, error) {
	var fields []*ReflectedField

	reflectedType := reflectValue(value)

	if cachedFields, ok := structFieldCache[reflectedType]; ok {
		return cachedFields, nil
	}

	for i := 0; i < reflectedType.NumField(); i++ {
		field := reflectedType.Field(i)
		if unicode.IsUpper(rune(field.Name[0])) { // only grab exported fields
			tagValue := field.Tag.Get(tagName)

			methodGroups, err := parser.Parse(tagValue)

			if err != nil {
				return nil, err
			}

			reflectedField := &ReflectedField{
				Index:        i,
				Name:         field.Name,
				MethodGroups: methodGroups,
			}

			fields = append(fields, reflectedField)
		}
	}

	structFieldCache[reflectedType] = fields

	return fields, nil
}

var (
	InvalidMethodError          = errors.New("Method does not exist.")
	InputParameterMismatchError = errors.New("Arguments does not match those of target function.")
	UnhandledCallError          = errors.New("Unhandled function call error.")
)

// Source: http://stackoverflow.com/questions/27673747/reflection-error-on-golang-too-few-arguments
// Note, this is really ugly/messy and should definitely be cleaned up.
func CallDynamicMethod(i interface{}, methodName string, args ...interface{}) ([]interface{}, error) {
	var ptr reflect.Value
	var value reflect.Value
	var finalMethod reflect.Value

	value = reflect.ValueOf(i)

	if value.Type().Kind() == reflect.Ptr {
		ptr = value
		value = ptr.Elem()
	} else {
		ptr = reflect.New(reflect.TypeOf(i))
		temp := ptr.Elem()
		temp.Set(value)
	}

	method := value.MethodByName(methodName)

	if method.IsValid() {
		finalMethod = method
	}

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

	return nil, UnhandledCallError
}
