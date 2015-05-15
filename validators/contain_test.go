package validators_test

import (
	"fmt"
	. "github.com/typerandom/validator/testing"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatContainValidatorFailsForInvalidOptions(t *testing.T) {
	dummy := 100

	ctx := NewTestContext(dummy)
	err := ContainValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error.")
	}

	err = ContainValidator(ctx, []interface{}{"123", "123"})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error.")
	}
}

func testThatContainValidatorSucceedsForExistingValue(t *testing.T, expect interface{}, dummy interface{}) {
	ctx := NewTestContext(dummy)
	opts := []interface{}{fmt.Sprintf("%v", expect)}

	if err := ContainValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got %s.", err)
	}
}

func testThatContainValidatorFailsForMissingValue(t *testing.T, expect interface{}, dummy interface{}) {
	ctx := NewTestContext(dummy)
	opts := []interface{}{fmt.Sprintf("%v", expect)}

	err := ContainValidator(ctx, opts)

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "contain.mustContainValue" {
		t.Fatalf("Expected must contain value error, got %s.", err)
	}
}

func TestThatContainValidatorSucceedsForExistingStringValue(t *testing.T) {
	testThatContainValidatorSucceedsForExistingValue(t, "2", "123")
	testThatContainValidatorSucceedsForExistingValue(t, "es", "test")
	testThatContainValidatorSucceedsForExistingValue(t, "test", "test")
}

func TestThatContainValidatorFailsForMissingStringValue(t *testing.T) {
	var dummy *string
	testThatContainValidatorFailsForMissingValue(t, "1", dummy)
	testThatContainValidatorFailsForMissingValue(t, "123", "")
	testThatContainValidatorFailsForMissingValue(t, "test1", "test")
}

func TestThatContainValidatorFailsForUnsupportedValueType(t *testing.T) {
	type Dummy struct{}

	ctx := NewTestContext(&Dummy{})
	err := ContainValidator(ctx, []interface{}{"123"})

	if err.Error() != "type.unsupported" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
