package validator

import (
	"github.com/typerandom/validator/core"
	"reflect"
)

func walkValidateArray(context *Context, normalized *core.NormalizedValue, parentField *core.ReflectedField) {
	valueType := reflect.ValueOf(normalized.Value)
	for i := 0; i < valueType.Len(); i++ {
		walkValidate(context, valueType.Index(i).Interface(), parentField)
	}
}

func walkValidateMap(context *Context, normalized *core.NormalizedValue, parentField *core.ReflectedField) {
	valueType := reflect.ValueOf(normalized.Value)
	for _, key := range valueType.MapKeys() {
		walkValidate(context, valueType.MapIndex(key).Interface(), parentField)
	}
}

func walkValidateStruct(context *Context, normalized *core.NormalizedValue, parentField *core.ReflectedField) {
	for _, field := range core.GetStructFields(normalized.Value, "validate") {
		normalizedFieldValue, err := core.NormalizeValue(field.Value, false)

		if err != nil {
			context.errors.Add(newValidatorError(field, nil, err))
			break
		}

		field.Parent = parentField

		context.setField(field)
		context.setSource(normalized.Value)
		context.setValue(normalizedFieldValue)

		for _, tag := range field.Tags {
			validate, err := getValidator(tag.Name)

			if err != nil {
				context.errors.Add(err)
				break
			}

			if err = validate(context, tag.Options); err != nil {
				context.errors.Add(newValidatorError(field, tag, err))
			}

			if context.StopValidate {
				break
			}
		}

		walkValidate(context, context.Value, field)
	}
}

func walkValidate(context *Context, value interface{}, parentField *core.ReflectedField) {
	var normalized *core.NormalizedValue

	if typedValue, ok := value.(*core.NormalizedValue); ok {
		normalized = typedValue
	} else {
		var err error
		normalized, err = core.NormalizeValue(value, false)
		if err != nil {
			context.errors.Add(err)
		}
	}

	switch normalized.OriginalKind {
	case reflect.Array, reflect.Slice:
		walkValidateArray(context, normalized, parentField)
	case reflect.Map:
		walkValidateMap(context, normalized, parentField)
	case reflect.Struct:
		walkValidateStruct(context, normalized, parentField)
	}
}
