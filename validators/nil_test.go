package validators

import (
	"errors"
	"github.com/typerandom/validator/core"
	"testing"
)

func TestThatItSucceedsForNilValue(t *testing.T) {
	var dummy *string

	ctx := core.NewTestContext(dummy)
	opts := []string{}

	if err := NilValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatItFailsForInvalidOptions(t *testing.T) {
	var dummy *string

	ctx := core.NewTestContext(dummy)
	opts := []string{"123"}

	err := NilValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.noneSupported" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatItFailsForNonNilValue(t *testing.T) {
	dummy := ""

	ctx := core.NewTestContext(&dummy)
	opts := []string{}

	err := NilValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "nil.isNotNil" {
		t.Fatal(errors.New("Expected is not nil error."))
	}
}

func TestThatItFailsForNonPointerValue(t *testing.T) {
	dummy := ""

	ctx := core.NewTestContext(dummy)
	opts := []string{}

	err := NilValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "nil.isNotNil" {
		t.Fatal(errors.New("Expected is not nil error."))
	}
}
