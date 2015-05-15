package validators_test

import (
	"errors"
	. "github.com/typerandom/validator/testing"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatLowerCaseValidatorFailsForInvalidOptions(t *testing.T) {
	var dummy string

	ctx := NewTestContext(dummy)
	opts := []interface{}{"123"}

	err := LowerCaseValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.noneSupported" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatLowerCaseValidatorSucceedsForEmptyString(t *testing.T) {
	dummy := ""

	ctx := NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatLowerCaseValidatorSucceedsForNilString(t *testing.T) {
	var dummy *string

	ctx := NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatLowerCaseValidatorFailsForUpperCaseString(t *testing.T) {
	dummy := "ABC"

	ctx := NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "lowerCase.mustBeLowerCase" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatLowerCaseValidatorFailsForMixedCaseString(t *testing.T) {
	dummy := "aBc"

	ctx := NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "lowerCase.mustBeLowerCase" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatLowerCaseValidatorSucceedsForStringWithoutCase(t *testing.T) {
	dummy := "123"

	ctx := NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatLowerCaseValidatorSucceedsForLowerCaseString(t *testing.T) {
	dummy := "abc"

	ctx := NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatLowerCaseValidatorFailsForUnsupportedValueType(t *testing.T) {
	type Dummy struct{}

	ctx := NewTestContext(&Dummy{})
	err := LowerCaseValidator(ctx, []interface{}{})

	if err.Error() != "type.unsupported" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
