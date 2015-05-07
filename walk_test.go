package validator_test

import (
	. "github.com/typerandom/validator"
	"reflect"
	"testing"
)

func testThatValidatorCanWalkItems(t *testing.T, items interface{}, numItems int) {
	errs := Validate(items)

	if len(errs.Items) != numItems {
		t.Fatalf("Expected %d errors, got %d.", numItems, len(errs.Items))
	}

	for _, err := range errs.Items {
		if err.Error() != "Value cannot be empty." {
			t.Fatalf("Expected validation error, got %s.", err)
		}
	}
}

type walkDummy struct {
	Value string `validate:"not_empty"`
}

func TestThatValidatorCanWalkSlice(t *testing.T) {
	dummies := []*walkDummy{&walkDummy{}, &walkDummy{}, &walkDummy{}, &walkDummy{}}
	testThatValidatorCanWalkItems(t, dummies, len(dummies))
}

func TestThatValidatorCanWalkArray(t *testing.T) {
	dummies := [...]*walkDummy{&walkDummy{}, &walkDummy{}, &walkDummy{}, &walkDummy{}}
	testThatValidatorCanWalkItems(t, dummies, len(dummies))
}

func TestThatValidatorCanWalkMap(t *testing.T) {
	dummies := map[string]*walkDummy{"a": &walkDummy{}, "b": &walkDummy{}, "c": &walkDummy{}, "d": &walkDummy{}}
	testThatValidatorCanWalkItems(t, dummies, len(dummies))
}

func TestThatValidatorCanWalkStruct(t *testing.T) {
	type Dummy struct {
		StringValue string  `validate:"not_empty"`
		IntValue    int     `validate:"not_empty"`
		FloatValue  float32 `validate:"not_empty"`
		UIntValue   uint    `validate:"not_empty"`
		PtrValue    *bool   `validate:"not_empty"`
	}

	dummy := Dummy{}
	errs := Validate(dummy)

	totalFields := reflect.TypeOf(dummy).NumField()

	if len(errs.Items) != totalFields {
		t.Fatalf("Expected %d errors, got %d.", totalFields, len(errs.Items))
	}

	expectedErrors := []string{
		"StringValue cannot be empty.",
		"IntValue cannot be empty.",
		"FloatValue cannot be empty.",
		"UIntValue cannot be empty.",
		"PtrValue cannot be empty.",
	}

	for i, err := range expectedErrors {
		if errs.Items[i].Error() != err {
			t.Fatalf("Expected error '%s', but got '%s'.", err, errs.Items[i].Error())
		}
	}
}

func testThatValidatorCannotWalkValue(t *testing.T, dummy interface{}, typeName string) {
	errs := Validate(dummy)

	if !errs.Any() {
		t.Fatalf("Expected error, but didn't get any.")
	}

	if len(errs.Items) != 1 {
		t.Fatalf("Expected 1 error, but got %d.", len(errs.Items))
	}

	expectedErr := "Unable to directly validate type '" + typeName + "'."

	if errs.First().Error() != expectedErr {
		t.Fatalf("Expected error '%s', but got '%s'.", expectedErr, errs.First())
	}
}

func TestThatValidatorCannotWalkString(t *testing.T) {
	testThatValidatorCannotWalkValue(t, "test", "string")
}

func TestThatValidatorCannotWalkInt(t *testing.T) {
	testThatValidatorCannotWalkValue(t, 123, "int64")
}

func TestThatValidatorCannotWalkFloat(t *testing.T) {
	testThatValidatorCannotWalkValue(t, 123.456, "float64")
}

func TestThatValidatorCannotWalkBool(t *testing.T) {
	testThatValidatorCannotWalkValue(t, false, "bool")
}

func TestThatValidatorCannotWalkInvalid(t *testing.T) {
	testThatValidatorCannotWalkValue(t, nil, "invalid")
}
