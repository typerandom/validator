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
