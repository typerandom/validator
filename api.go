package main

import (
	"reflect"
)

func walkValidateArray(context *ValidatorContext, normalized *normalizedValue, parentField *reflectedField) {
	valueType := reflect.ValueOf(normalized.Value)
	for i := 0; i < valueType.Len(); i++ {
		walkValidate(context, valueType.Index(i).Interface(), parentField)
	}
}

func walkValidateMap(context *ValidatorContext, normalized *normalizedValue, parentField *reflectedField) {
	valueType := reflect.ValueOf(normalized.Value)
	for _, key := range valueType.MapKeys() {
		walkValidate(context, valueType.MapIndex(key).Interface(), parentField)
	}
}

func walkValidateStruct(context *ValidatorContext, normalized *normalizedValue, parentField *reflectedField) {
	for _, field := range getFields(normalized.Value, "validate") {
		normalizedFieldValue, err := normalizeValue(field.Value, false)

		if err != nil {
			context.Errors.Add(newValidatorError(field, nil, err))
			break
		}

		field.Parent = parentField

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

		walkValidate(context, context.Value, field)
	}
}

func walkValidate(context *ValidatorContext, value interface{}, parentField *reflectedField) {
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
		walkValidateArray(context, normalized, parentField)
	case reflect.Map:
		walkValidateMap(context, normalized, parentField)
	case reflect.Struct:
		context.SetSource(normalized.Value)
		walkValidateStruct(context, normalized, parentField)
	}
}

func Register(name string, validator ValidatorFilter) {
	registerValidator(name, validator)
}

func Validate(value interface{}) *Errors {
	context := &ValidatorContext{
		Errors: NewErrors(),
	}

	walkValidate(context, value, nil)

	return context.Errors
}
