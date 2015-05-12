package validator

import (
	"errors"
	"github.com/typerandom/validator/core"
	"reflect"
)

func canWalk(value reflect.Kind) bool {
	switch value {
	case reflect.Ptr, reflect.Array, reflect.Slice, reflect.Map, reflect.Struct:
		return true
	default:
		return false
	}
}

func walkValidateArray(context *context, normalized *core.NormalizedValue, parentField *core.ReflectedField) {
	valueType := reflect.ValueOf(normalized.Value)
	for i := 0; i < valueType.Len(); i++ {
		value := valueType.Index(i)
		if canWalk(value.Kind()) {
			walkValidate(context, value.Interface(), parentField)
		}
	}
}

func walkValidateMap(context *context, normalized *core.NormalizedValue, parentField *core.ReflectedField) {
	valueType := reflect.ValueOf(normalized.Value)
	for _, key := range valueType.MapKeys() {
		value := valueType.MapIndex(key)
		if canWalk(value.Kind()) {
			walkValidate(context, value.Interface(), parentField)
		}
	}
}

func walkValidateStruct(context *context, normalized *core.NormalizedValue, parentField *core.ReflectedField) {
	fields, err := core.GetStructFields(normalized.Value, "validate")

	if err != nil {
		context.errors.Add(err)
		return
	}

	sourceStruct := reflect.Indirect(reflect.ValueOf(normalized.Value))

	for _, field := range fields {
		fieldValue := field.GetValue(sourceStruct)

		normalizedFieldValue, err := core.Normalize(fieldValue)

		if err != nil {
			context.errors.Add(core.NewValidatorError(field, nil, err))
			continue
		}

		field.Parent = parentField

		context.setField(field)
		context.setSource(normalized.Value)
		context.setValue(normalizedFieldValue)

		var mostRecentErrors *core.Errors

		for _, methods := range field.MethodGroups {
			var errors *core.Errors

			for _, method := range methods {
				validate, err := context.validator.registry.Get(method.Name)

				if err != nil {
					context.errors.Add(err)
					return
				}

				if err = validate(context, method.Arguments); err != nil {
					if errors == nil {
						errors = core.NewErrors()
					}
					errors.Add(core.NewValidatorError(field, method, err))
				}
			}

			mostRecentErrors = errors

			if errors == nil {
				break
			}
		}

		if mostRecentErrors != nil {
			context.errors.AddMany(mostRecentErrors)
		}

		if canWalk(normalizedFieldValue.OriginalKind) {
			walkValidate(context, normalizedFieldValue, field)
		}
	}
}

func walkValidate(context *context, value interface{}, parentField *core.ReflectedField) {
	var normalized *core.NormalizedValue

	if typedValue, ok := value.(*core.NormalizedValue); ok {
		normalized = typedValue
	} else {
		var err error
		normalized, err = core.Normalize(value)
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
	default:
		context.errors.Add(errors.New("Unable to directly validate type '" + normalized.OriginalKind.String() + "'."))
	}
}
