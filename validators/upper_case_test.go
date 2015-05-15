package validators_test

import (
	"errors"
	. "github.com/typerandom/validator/testing"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatUpperCaseValidatorFailsForInvalidOptions(t *testing.T) {
	var dummy string

	ctx := NewTestContext(dummy)
	opts := []interface{}{"123"}

	err := UpperCaseValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.noneSupported" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatUpperCaseValidatorSucceedsForEmptyString(t *testing.T) {
	dummy := ""

	ctx := NewTestContext(dummy)
	err := UpperCaseValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatUpperCaseValidatorSucceedsForNilString(t *testing.T) {
	var dummy *string

	ctx := NewTestContext(dummy)
	err := UpperCaseValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatUpperCaseValidatorFailsForLowerCaseString(t *testing.T) {
	dummy := "abc"

	ctx := NewTestContext(dummy)
	err := UpperCaseValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "upperCase.mustBeUpperCase" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatUpperCaseValidatorFailsForMixedCaseString(t *testing.T) {
	dummy := "aBc"

	ctx := NewTestContext(dummy)
	err := UpperCaseValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "upperCase.mustBeUpperCase" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatUpperCaseValidatorSucceedsForStringWithoutCase(t *testing.T) {
	dummy := "123"

	ctx := NewTestContext(dummy)
	err := UpperCaseValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatUpperCaseValidatorSucceedsForUpperCaseString(t *testing.T) {
	dummy := "ABC"

	ctx := NewTestContext(dummy)
	err := UpperCaseValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatUpperCaseValidatorFailsForUnsupportedValueType(t *testing.T) {
	type Dummy struct{}

	ctx := NewTestContext(&Dummy{})
	err := UpperCaseValidator(ctx, []interface{}{})

	if err.Error() != "type.unsupported" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
