package main

import (
	"reflect"
)

func Register(name string, validator ValidatorFilter) {
	registerValidator(name, validator)
}

func Validate(value interface{}) *Errors {
	context := &ValidatorContext{
		Errors: NewErrors(),
	}

	walkValidate(context, value)

	return context.Errors
}

func walkValidateArray(context *ValidatorContext, normalized *normalizedValue) {
	valueType := reflect.ValueOf(normalized.Value)
	for i := 0; i < valueType.Len(); i++ {
		walkValidate(context, valueType.Index(i).Interface())
	}
}

func walkValidateMap(context *ValidatorContext, normalized *normalizedValue) {
	valueType := reflect.ValueOf(normalized.Value)
	for _, key := range valueType.MapKeys() {
		walkValidate(context, valueType.MapIndex(key).Interface())
	}
}

func walkValidateStruct(context *ValidatorContext, normalized *normalizedValue) {
	for _, field := range getFields(normalized.Value, "validate") {
		normalizedFieldValue, err := normalizeValue(field.Value, false)

		if err != nil {
			context.Errors.Add(newValidatorError(field, nil, err))
			break
		}

		context.SetField(field)
		context.SetValue(normalizedFieldValue)

		for _, tag := range field.Tags {
			validate, err := getValidator(tag.Name)

			if err != nil {
				context.Errors.Add(err)
				break
			}

			if err = validate(context, tag.Options); err != nil {
				context.Errors.Add(newValidatorError(field, tag, err))
			}

			if context.StopValidate {
				break
			}
		}

		walkValidate(context, context.Value)
	}
}

func walkValidate(context *ValidatorContext, value interface{}) {
	var normalized *normalizedValue

	if typedValue, ok := value.(*normalizedValue); ok {
		normalized = typedValue
	} else {
		var err error
		normalized, err = normalizeValue(value, false)
		if err != nil {
			context.Errors.Add(err)
		}
	}

	switch normalized.OriginalKind {
	case reflect.Array, reflect.Slice:
		walkValidateArray(context, normalized)
	case reflect.Map:
		walkValidateMap(context, normalized)
	case reflect.Struct:
		context.SetParent(normalized.Value)
		walkValidateStruct(context, normalized)
	}
}
