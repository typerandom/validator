package validators_test

import (
	"github.com/typerandom/validator/core"
	. "github.com/typerandom/validator/validators"
	"testing"
	"time"
)

func TestThatTimeValidatorFailsForInvalidOptions(t *testing.T) {
	dummy := "100"

	ctx := core.NewTestContext(dummy)
	err := TimeValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error, got %s.", err)
	}

	err = TimeValidator(ctx, []interface{}{"123", "123"})

	if err == nil {
		t.Fatal("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error, got %s.", err)
	}
}

func TestThatTimeValidatorSucceedsForValidStringTimeValue(t *testing.T) {
	dummy := "2013-06-05T14:10:43.678Z"

	ctx := core.NewTestContext(dummy)
	opts := []interface{}{time.RFC3339Nano}

	if err := TimeValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatTimeValidatorFailsForInvalidNilStringValue(t *testing.T) {
	var dummy *string

	ctx := core.NewTestContext(dummy)
	opts := []interface{}{time.RFC3339Nano}

	err := TimeValidator(ctx, opts)

	if err == nil {
		t.Fatalf("Expected error, got none.")
	}

	if err.Error() != "time.mustBeValid" {
		t.Fatalf("Expected time must be valid error, got %s.", err)
	}
}

func TestThatTimeValidatorFailsForInvalidStringTimeValue(t *testing.T) {
	dummy := "2013-06-05INVALID14:10:43"

	ctx := core.NewTestContext(dummy)
	opts := []interface{}{time.RFC3339Nano}

	err := TimeValidator(ctx, opts)

	if err == nil {
		t.Fatalf("Expected error, got none.")
	}

	if err.Error() != "time.mustBeValid" {
		t.Fatalf("Expected time must be valid error, got %s.", err)
	}
}

func TestThatTimeValidatorSucceedsForTimeValue(t *testing.T) {
	dummy := time.Now()

	ctx := core.NewTestContext(dummy)
	opts := []interface{}{}

	if err := TimeValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatTimeValidatorFailsForUnsupportedValueType(t *testing.T) {
	type Dummy struct{}

	ctx := core.NewTestContext(&Dummy{})
	err := TimeValidator(ctx, []interface{}{})

	if err.Error() != "type.unsupported" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
