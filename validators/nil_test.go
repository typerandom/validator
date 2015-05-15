package validators_test

import (
	"errors"
	. "github.com/typerandom/validator/testing"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatNilValidatorFailsForInvalidOptions(t *testing.T) {
	var dummy *string

	ctx := NewTestContext(dummy)
	opts := []interface{}{"123"}

	err := NilValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.noneSupported" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatNilValidatorSucceedsForNilValue(t *testing.T) {
	var dummy *string

	ctx := NewTestContext(dummy)
	opts := []interface{}{}

	if err := NilValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatNilValidatorFailsForNonNilValue(t *testing.T) {
	dummy := ""

	ctx := NewTestContext(&dummy)
	opts := []interface{}{}

	err := NilValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "nil.isNotNil" {
		t.Fatal(errors.New("Expected is not nil error."))
	}
}

func TestThatNilValidatorFailsForNonPointerValue(t *testing.T) {
	dummy := ""

	ctx := NewTestContext(dummy)
	opts := []interface{}{}

	err := NilValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "nil.isNotNil" {
		t.Fatal(errors.New("Expected is not nil error."))
	}
}
