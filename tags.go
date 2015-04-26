package main

import (
	"bytes"
	"reflect"
	"strings"
)

const (
	STATE_NAME              = 0
	STATE_OPTION            = 1
	STATE_OPTION_ESCAPE     = 2
	STATE_OPTION_WHITESPACE = 3
)

type tag struct {
	Name    string
	Options []string
}

func (this *tag) String() string {
	serializedOptions := strings.Join(this.Options, ", ")

	if len(serializedOptions) == 0 {
		serializedOptions = "(none)"
	} else {
		serializedOptions = "'" + serializedOptions + "'"
	}

	return "{ name: '" + this.Name + "', options: " + serializedOptions + " }"
}

func parseTag(rawTag string) []*tag {
	var tags []*tag

	var buffer bytes.Buffer
	var currentTag *tag
	var currentState int

	for _, char := range rawTag {
		switch {
		case currentState == STATE_OPTION_ESCAPE:
			buffer.WriteRune(char)
			currentState = STATE_OPTION
		case currentState == STATE_OPTION && char == '\\':
			currentState = STATE_OPTION_ESCAPE
		case char == '(':
			currentState = STATE_OPTION
			currentTag.Name = buffer.String()
			buffer.Reset()
		case currentState == STATE_OPTION && char == ')':
			currentState = STATE_NAME
			if buffer.Len() > 0 {
				currentTag.Options = append(currentTag.Options, buffer.String())
				buffer.Reset()
			}
		case currentState == STATE_OPTION_WHITESPACE:
			if char != ' ' && char != '	' {
				currentState = STATE_OPTION
				buffer.WriteRune(char)
			}
		case currentState == STATE_OPTION && char == ',':
			if buffer.Len() > 0 {
				currentState = STATE_OPTION_WHITESPACE
				currentTag.Options = append(currentTag.Options, buffer.String())
				buffer.Reset()
			}
		case currentState == STATE_NAME && char == ',':
			if buffer.Len() > 0 {
				currentTag.Name = buffer.String()
				buffer.Reset()
			}

			if len(currentTag.Name) > 0 {
				tags = append(tags, currentTag)
			}

			currentTag = &tag{}
		default:
			if currentTag == nil {
				currentTag = &tag{}
			}
			buffer.WriteRune(char)
		}
	}

	if buffer.Len() > 0 {
		currentTag.Name = buffer.String()
	}

	if currentTag != nil && len(currentTag.Name) > 0 {
		tags = append(tags, currentTag)
	}

	return tags
}

type reflectedField struct {
	Parent *reflectedField
	Name   string
	Value  interface{}
	Tags   []*tag
}

func newReflectedField(name string, value interface{}, tag string) *reflectedField {
	return &reflectedField{
		Name:  name,
		Value: value,
		Tags:  parseTag(tag),
	}
}

func (this *reflectedField) FullName() string {
	fullName := this.Name
	parent := this.Parent

	for parent != nil {
		if len(parent.Name) > 0 {
			fullName = parent.Name + "." + fullName
		}
		parent = parent.Parent
	}

	return fullName
}

func reflectValue(value interface{}) (reflect.Type, reflect.Value) {
	reflectValue := reflect.ValueOf(value)
	valueType := reflectValue.Type()

	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}

	return valueType, reflect.Indirect(reflectValue)
}

func getFields(value interface{}, tagName string) []*reflectedField {
	var fields []*reflectedField

	valueType, reflectedValue := reflectValue(value)

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		tagValue := field.Tag.Get(tagName)
		fields = append(fields, newReflectedField(field.Name, reflectedValue.Field(i).Interface(), tagValue))
	}

	return fields
}
