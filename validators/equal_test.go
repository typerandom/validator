package validators_test

import (
	"fmt"
	. "github.com/typerandom/validator/testing"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatEqualValidatorFailsForInvalidOptions(t *testing.T) {
	dummy := 100

	ctx := NewTestContext(dummy)
	err := EqualValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error.")
	}

	err = EqualValidator(ctx, []interface{}{"123", "123"})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error.")
	}
}

func testThatEqualValidatorSucceedsForEqualValue(t *testing.T, expect interface{}, dummy interface{}) {
	ctx := NewTestContext(dummy)
	opts := []interface{}{fmt.Sprintf("%v", expect)}

	if err := EqualValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func testThatEqualValidatorFailsForNonEqualValue(t *testing.T, expect interface{}, dummy interface{}) {
	ctx := NewTestContext(dummy)
	opts := []interface{}{fmt.Sprintf("%v", expect)}

	err := EqualValidator(ctx, opts)

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "equal.mustEqualValue" {
		t.Fatalf("Expected must equal value error.")
	}
}

func TestThatEqualValidatorSucceedsForEqualStringValue(t *testing.T) {
	testThatEqualValidatorSucceedsForEqualValue(t, "", "")
	testThatEqualValidatorSucceedsForEqualValue(t, "test", "test")
}

func TestThatEqualValidatorFailsForNonEqualStringValue(t *testing.T) {
	var dummy *string
	testThatEqualValidatorFailsForNonEqualValue(t, "", dummy)
	testThatEqualValidatorFailsForNonEqualValue(t, "test", "test1")
}

func TestThatEqualValidatorSucceedsForEqualIntValue(t *testing.T) {
	testThatEqualValidatorSucceedsForEqualValue(t, 0, 0)
	testThatEqualValidatorSucceedsForEqualValue(t, 1234, 1234)
}

func TestThatEqualValidatorFailsForNonEqualIntValue(t *testing.T) {
	var dummy *int64
	testThatEqualValidatorFailsForNonEqualValue(t, 0, dummy)
	testThatEqualValidatorFailsForNonEqualValue(t, 1234, 12345)
}

func TestThatEqualValidatorSucceedsForEqualFloatValue(t *testing.T) {
	testThatEqualValidatorSucceedsForEqualValue(t, 0.0, 0.0)
	testThatEqualValidatorSucceedsForEqualValue(t, 1.234, 1.234)
}

func TestThatEqualValidatorFailsForNonEqualFloatValue(t *testing.T) {
	var dummy *float64
	testThatEqualValidatorFailsForNonEqualValue(t, 0.0, dummy)
	testThatEqualValidatorFailsForNonEqualValue(t, 1.234, 1.2345)
}

func TestThatEqualValidatorSucceedsForEqualBoolValue(t *testing.T) {
	testThatEqualValidatorSucceedsForEqualValue(t, false, false)
	testThatEqualValidatorSucceedsForEqualValue(t, true, true)
}

func TestThatEqualValidatorFailsForNonEqualBoolValue(t *testing.T) {
	var dummy *bool
	testThatEqualValidatorFailsForNonEqualValue(t, false, dummy)
	testThatEqualValidatorFailsForNonEqualValue(t, false, true)
}

func TestThatEqualValidatorFailsForUnsupportedValueType(t *testing.T) {
	type Dummy struct{}

	ctx := NewTestContext(&Dummy{})
	err := EqualValidator(ctx, []interface{}{"123"})

	if err.Error() != "type.unsupported" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
