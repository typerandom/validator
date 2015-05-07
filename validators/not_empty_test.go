package validators_test

import (
	"errors"
	"github.com/typerandom/validator/core"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatNotEmptyValidatorSucceedsForInvalidOptions(t *testing.T) {
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

func testThatNotEmptyValidatorFailsForEmptyValue(t *testing.T, dummy interface{}) {
	ctx := core.NewTestContext(dummy)
	opts := []string{}

	if err := EmptyValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func testThatNotEmptyValidatorSucceedsNonEmptyValue(t *testing.T, dummy interface{}) {
	ctx := core.NewTestContext(&dummy)
	err := EmptyValidator(ctx, []string{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "empty.isNotEmpty" {
		t.Fatal(errors.New("Expected is not nil error."))
	}
}

func TestThatNotEmptyValidatorFailsForEmptyStringValue(t *testing.T) {
	testThatNotEmptyValidatorFailsForEmptyValue(t, "")
}

func TestThatNotEmptyValidatorFailsForStringNilValue(t *testing.T) {
	var dummy *string
	testThatNotEmptyValidatorFailsForEmptyValue(t, dummy)
}

func TestThatNotEmptyValidatorSucceedsForNonEmptyStringValue(t *testing.T) {
	dummy := "test"
	testThatNotEmptyValidatorSucceedsNonEmptyValue(t, dummy)
}

func TestThatNotEmptyValidatorFailsForZeroIntValue(t *testing.T) {
	testThatNotEmptyValidatorFailsForEmptyValue(t, 0)
}

func TestThatNotEmptyValidatorFailsForIntNilValue(t *testing.T) {
	var dummy *int64
	testThatNotEmptyValidatorFailsForEmptyValue(t, dummy)
}

func TestThatNotEmptyValidatorSucceedsForNonZeroIntValue(t *testing.T) {
	dummy := 123
	testThatNotEmptyValidatorSucceedsNonEmptyValue(t, dummy)
}

func TestThatNotEmptyValidatorFailsForZeroFloatValue(t *testing.T) {
	testThatNotEmptyValidatorFailsForEmptyValue(t, 0.0)
}

func TestThatNotEmptyValidatorFailsForFloatNilValue(t *testing.T) {
	var dummy *float64
	testThatNotEmptyValidatorFailsForEmptyValue(t, dummy)
}

func TestThatNotEmptyValidatorSucceedsForNonZeroFloatValue(t *testing.T) {
	dummy := 1.23
	testThatNotEmptyValidatorSucceedsNonEmptyValue(t, dummy)
}

func TestThatNotEmptyValidatorFailsForFalseBoolValue(t *testing.T) {
	testThatNotEmptyValidatorFailsForEmptyValue(t, false)
}

func TestThatNotEmptyValidatorSucceedsForTrueBoolValue(t *testing.T) {
	testThatNotEmptyValidatorSucceedsNonEmptyValue(t, true)
}

func TestThatNotEmptyValidatorFailsForEmptySliceValue(t *testing.T) {
	testThatNotEmptyValidatorFailsForEmptyValue(t, []string{})
}

func TestThatNotEmptyValidatorSucceedsForNonEmptySliceValue(t *testing.T) {
	testThatNotEmptyValidatorSucceedsNonEmptyValue(t, []string{"abc"})
}

func TestThatNotEmptyValidatorFailsForEmptyMapValue(t *testing.T) {
	testThatNotEmptyValidatorFailsForEmptyValue(t, map[string]string{})
}

func TestThatNotEmptyValidatorSucceedsForNonEmptyMapValue(t *testing.T) {
	testThatNotEmptyValidatorSucceedsNonEmptyValue(t, map[string]string{"abc": "abc"})
}

func TestThatNotEmptyValidatorSucceedsForUnhandledType(t *testing.T) {
	type Dummy struct{}

	ctx := core.NewTestContext(&Dummy{})
	err := EmptyValidator(ctx, []string{})

	if err.Error() != "empty.isNotEmpty" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
