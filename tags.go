package main

import (
	"bytes"
	"reflect"
)

const (
	STATE_NAME          = 0
	STATE_OPTION        = 1
	STATE_OPTION_ESCAPE = 2
)

type tag struct {
	Name    string
	Options string
}

func (this *tag) String() string {
	serializedOptions := this.Options

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
		/*case currentState == STATE_OPTION_ESCAPE:
			buffer += string(char)
			currentState = STATE_OPTION
		case currentState == STATE_OPTION && char == '\\':
			currentState = STATE_OPTION_ESCAPE*/
		case char == '(':
			currentState = STATE_OPTION
			currentTag.Name = buffer.String()
			buffer.Reset()
		case currentState == STATE_OPTION && char == ')':
			currentState = STATE_NAME
			currentTag.Options = buffer.String()
			buffer.Reset()
		case currentState == STATE_NAME && char == ',':
			currentState = STATE_NAME

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

type taggedField struct {
	Name  string
	Value interface{}
	Tags  []*tag
}

func newTaggedField(name string, value interface{}, tag string) *taggedField {
	return &taggedField{
		Name:  name,
		Value: value,
		Tags:  parseTag(tag),
	}
}

type ValueContext struct {
	Value interface{}
	Type  reflect.Type
	IsNil bool
}

func getTaggedFields(value interface{}, tagName string) []*taggedField {
	var fields []*taggedField

	val := reflect.ValueOf(value)
	valueType := val.Type()

	if valueType.Kind() == reflect.Ptr {
		valueType = valueType.Elem()
	}

	for i := 0; i < valueType.NumField(); i++ {
		field := valueType.Field(i)
		tagValue := field.Tag.Get(tagName)
		if len(tagValue) > 0 {
			fieldValue := reflect.Indirect(val).Field(i).Interface()
			fields = append(fields, newTaggedField(field.Name, fieldValue, field.Tag.Get(tagName)))
		}
	}

	return fields
}