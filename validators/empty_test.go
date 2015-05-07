package validators_test

import (
	"errors"
	"github.com/typerandom/validator/core"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatEmptyValidatorFailsForInvalidOptions(t *testing.T) {
	var dummy *string

	ctx := core.NewTestContext(dummy)
	opts := []string{"123"}

	err := EmptyValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.noneSupported" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func testThatEmptyValidatorSucceedsForEmptyValue(t *testing.T, dummy interface{}) {
	ctx := core.NewTestContext(dummy)
	opts := []string{}

	if err := EmptyValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func testThatEmptyValidatorFailsForNonEmptyValue(t *testing.T, dummy interface{}) {
	ctx := core.NewTestContext(&dummy)
	err := EmptyValidator(ctx, []string{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "empty.isNotEmpty" {
		t.Fatal(errors.New("Expected is not nil error."))
	}
}

func TestThatEmptyValidatorSucceedsForEmptyStringValue(t *testing.T) {
	testThatEmptyValidatorSucceedsForEmptyValue(t, "")
}

func TestThatEmptyValidatorSucceedsForStringNilValue(t *testing.T) {
	var dummy *string
	testThatEmptyValidatorSucceedsForEmptyValue(t, dummy)
}

func TestThatEmptyValidatorFailsForNonEmptyStringValue(t *testing.T) {
	dummy := "test"
	testThatEmptyValidatorFailsForNonEmptyValue(t, dummy)
}

func TestThatEmptyValidatorSucceedsForZeroIntValue(t *testing.T) {
	testThatEmptyValidatorSucceedsForEmptyValue(t, 0)
}

func TestThatEmptyValidatorSucceedsForIntNilValue(t *testing.T) {
	var dummy *int64
	testThatEmptyValidatorSucceedsForEmptyValue(t, dummy)
}

func TestThatEmptyValidatorFailsForNonZeroIntValue(t *testing.T) {
	dummy := 123
	testThatEmptyValidatorFailsForNonEmptyValue(t, dummy)
}

func TestThatEmptyValidatorSucceedsForZeroFloatValue(t *testing.T) {
	testThatEmptyValidatorSucceedsForEmptyValue(t, 0.0)
}

func TestThatEmptyValidatorSucceedsForFloatNilValue(t *testing.T) {
	var dummy *float64
	testThatEmptyValidatorSucceedsForEmptyValue(t, dummy)
}

func TestThatEmptyValidatorFailsForNonZeroFloatValue(t *testing.T) {
	dummy := 1.23
	testThatEmptyValidatorFailsForNonEmptyValue(t, dummy)
}

func TestThatEmptyValidatorSucceedsForFalseBoolValue(t *testing.T) {
	testThatEmptyValidatorSucceedsForEmptyValue(t, false)
}

func TestThatEmptyValidatorFailsForTrueBoolValue(t *testing.T) {
	testThatEmptyValidatorFailsForNonEmptyValue(t, true)
}

func TestThatEmptyValidatorSucceedsForEmptySliceValue(t *testing.T) {
	testThatEmptyValidatorSucceedsForEmptyValue(t, []string{})
}

func TestThatEmptyValidatorFailsForNonEmptySliceValue(t *testing.T) {
	testThatEmptyValidatorFailsForNonEmptyValue(t, []string{"abc"})
}

func TestThatEmptyValidatorSucceedsForEmptyMapValue(t *testing.T) {
	testThatEmptyValidatorSucceedsForEmptyValue(t, map[string]string{})
}

func TestThatEmptyValidatorFailsForNonEmptyMapValue(t *testing.T) {
	testThatEmptyValidatorFailsForNonEmptyValue(t, map[string]string{"abc": "abc"})
}

func TestThatEmptyValidatorFailsForUnhandledType(t *testing.T) {
	type Dummy struct{}

	ctx := core.NewTestContext(&Dummy{})
	err := EmptyValidator(ctx, []string{})

	if err.Error() != "empty.isNotEmpty" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
