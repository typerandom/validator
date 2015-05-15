package core_test

import (
	. "github.com/typerandom/validator/core"
	"testing"
)

func TestThatSpecificStructFieldsCanBeReflected(t *testing.T) {
	type Foo struct {
		ValueA string `test:"abc"`
		ValueB int    `test:"def" name:"custom_value"`
	}

	value := &Foo{
		ValueA: "abc",
		ValueB: 123,
	}

	fields, err := GetStructFields(value, "test", "name")

	if err != nil {
		t.Fatalf("Didn't expect an error, but got '%s'.", err)
	}

	if len(fields) != 2 {
		t.Fatalf("Expected 2 fields, but got %d.", len(fields))
	}

	firstField := fields[0]

	if firstField.Index != 0 {
		t.Fatalf("Expected index of first field to be '0', but got '%d'.", firstField.Index)
	}

	if firstField.DisplayName != "ValueA" {
		t.Fatalf("Expected display name of first field to be 'custom_value', but got '%s'.", firstField.DisplayName)
	}

	if firstField.FullName() != "ValueA" {
		t.Fatalf("Expected full name of first field to be 'ValueA', but got '%s'.", firstField.FullName())
	}

	if firstField.Name != "ValueA" {
		t.Fatalf("Expected name of first field to be 'ValueA', but got '%s'.", firstField.Name)
	}

	if len(firstField.MethodGroups) == 0 {
		t.Fatalf("Expected there to be one method groups, but got that there wasn't any.")
	} else {
		if len(firstField.MethodGroups) > 1 || len(firstField.MethodGroups[0]) != 1 {
			t.Fatalf("Expected there to be 1 method groups with 1 methods of first field, but there was '%d' and '%d'.", len(firstField.MethodGroups), len(firstField.MethodGroups[0]))
		}
		if firstField.MethodGroups[0][0].Name != "abc" {
			t.Fatalf("Exepected first method of method group to be 'abc', but got '%s'.", firstField.MethodGroups[0][0].Name)
		}
	}

	secondField := fields[1]

	if secondField.Index != 1 {
		t.Fatalf("Expected index of first field to be '1', but got '%d'.", secondField.Index)
	}

	if secondField.DisplayName != "custom_value" {
		t.Fatalf("Expected display name of first field to be 'custom_value', but got '%s'.", secondField.DisplayName)
	}

	if secondField.FullName() != "ValueB" {
		t.Fatalf("Expected full name of first field to be 'ValueB', but got '%s'.", secondField.FullName())
	}

	if secondField.Name != "ValueB" {
		t.Fatalf("Expected name of first field to be 'ValueB', but got '%s'.", secondField.Name)
	}

	if len(secondField.MethodGroups) == 0 {
		t.Fatalf("Expected there to be one method groups, but got that there wasn't any.")
	} else {
		if len(secondField.MethodGroups) > 1 || len(secondField.MethodGroups[0]) != 1 {
			t.Fatalf("Expected there to be 1 method groups with 1 methods of first field, but there was '%d' and '%d'.", len(secondField.MethodGroups), len(secondField.MethodGroups[0]))
		}
		if secondField.MethodGroups[0][0].Name != "def" {
			t.Fatalf("Exepected first method of method group to be 'def', but got '%s'.", secondField.MethodGroups[0][0].Name)
		}
	}
}
