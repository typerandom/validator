// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Originally from https://golang.org/src/encoding/json/tags.go?m=text

package main

import (
	"reflect"
)

type taggedField struct {
	Name  string
	Value interface{}
	Tag   string
}

func newTaggedField(name string, value interface{}, tag string) *taggedField {
	return &taggedField{
		Name:  name,
		Value: value,
		Tag:   tag,
	}
}

type tag struct {
	name    string
	options string
}

func (this *tag) String() string {
	return "{ name=" + this.name + ", options=" + this.options + " },"
}

const (
	STATE_NAME          = 0
	STATE_OPTION        = 1
	STATE_OPTION_ESCAPE = 2
)

func parseTag(rawTag string) []*tag {
	var tags []*tag

	var buffer string
	var currentTag *tag
	currentState := STATE_NAME

	for _, char := range rawTag {
		switch {
		/*case currentState == STATE_OPTION_ESCAPE:
			buffer += string(char)
			currentState = STATE_OPTION
		case currentState == STATE_OPTION && char == '\\':
			currentState = STATE_OPTION_ESCAPE*/
		case char == '(':
			currentState = STATE_OPTION
			currentTag.name = buffer
			buffer = ""
		case currentState == STATE_OPTION && char == ')':
			currentState = STATE_NAME
			currentTag.options = buffer
			buffer = ""
		case currentState == STATE_NAME && char == ',':
			currentState = STATE_NAME

			if len(buffer) > 0 {
				currentTag.name = buffer
				buffer = ""
			}

			if len(currentTag.name) > 0 {
				tags = append(tags, currentTag)
			}

			currentTag = &tag{}
		default:
			if currentTag == nil {
				currentTag = &tag{}
			}
			buffer += string(char)
		}
	}

	if len(buffer) > 0 {
		currentTag.name = buffer
	}

	if currentTag != nil && len(currentTag.name) > 0 {
		tags = append(tags, currentTag)
	}

	return tags
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
