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

func walkValidate(context *ValidatorContext, value interface{}, parentField *reflectedField) {
	var normalized *normalizedValue

	if typedValue, ok := value.(*normalizedValue); ok {
		normalized = typedValue
	} else {
		var err error
		normalized, err = normalizeValue(value, false)
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

var globalInitialized bool
var globalLocale *locale

func assertGlobalInit() {
	if !globalInitialized {
		globalInitialized = true
		if globalLocale == nil {
			globalLocale = newLocale()
			registerDefaultLocale(globalLocale)
		}
	}
}

func registerDefaultLocale(lc *locale) {
	lc.Set("arguments.invalid", "Unable to parse '{validator}' validator options for field '{field}'.")
	lc.Set("arguments.noneSupported", "Validator '{validator}' on field '{field}' does not support any arguments.")
	lc.Set("arguments.singleRequired", "Validator '{validator}' on field '{field}' requires a single argument.")
	lc.Set("arguments.oneOrMoreRequired", "Validator '{validator}' on field '{field}' requires at least one argument.")
	lc.Set("notEmpty.cannotBeEmpty", "{field} cannot be empty.")
	lc.Set("min.cannotBeShorterThan", "{field} cannot be shorter than %d characters.")
	lc.Set("min.cannotBeLessThan", "{field} cannot be less than %d.")
	lc.Set("min.cannotContainLessItemsThan", "{field} cannot contain less than %d items.")
	lc.Set("min.cannotContainLessKeysThan", "{field} cannot contain less than %d keys.")
	lc.Set("max.cannotBeLongerThan", "{field} is longer than %d characters.")
	lc.Set("max.cannotBeGreaterThan", "{field} cannot be greater than %d.")
	lc.Set("max.cannotContainMoreItemsThan", "{field} cannot contain more than %d items.")
	lc.Set("max.cannotContainMoreKeysThan", "{field} cannot contain more than %d keys.")
	lc.Set("lowerCase.mustBeLowerCase", "{field} must be in lower case.")
	lc.Set("upperCase.mustBeUpperCase", "{field} must be in upper case.")
	lc.Set("contain.mustContainValues", "{field} must contain one of the following values '%s'.")
	lc.Set("equal.mustEqualValues", "{field} must equal one of the following values '%s'.")
	lc.Set("regexp.mustMatchPattern", "{field} must match pattern '%s'.")
	lc.Set("numeric.mustBeNumeric", "{field} must be numeric.")
	lc.Set("time.mustBeValid", "{field} must be a valid time.")
}

func LoadLocale(jsonPath string) error {
	assertGlobalInit()
	return globalLocale.LoadJson(jsonPath)
}

func Register(name string, validator ValidatorFilter) {
	registerValidator(name, validator)
}

func Validate(value interface{}) *Errors {
	assertGlobalInit()

	context := &ValidatorContext{
		errors: &Errors{},
		locale: globalLocale,
	}

	walkValidate(context, value, nil)

	return context.errors
}
