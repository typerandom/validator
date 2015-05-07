package validators_test

import (
	"github.com/typerandom/validator/core"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatRegexpValidatorFailsForInvalidOptions(t *testing.T) {
	dummy := 100

	ctx := core.NewTestContext(dummy)
	err := RegexpValidator(ctx, []string{})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error.")
	}

	err = RegexpValidator(ctx, []string{"123", "123"})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error.")
	}
}

func TestThatRegexpValidatorSucceedsForMatchingStringValue(t *testing.T) {
	dummy := "1test."

	ctx := core.NewTestContext(dummy)
	opts := []string{"^\\dtest\\.$"}

	if err := RegexpValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatRegexpValidatorFailsForNonMatchingStringValue(t *testing.T) {
	dummy := "1test"

	ctx := core.NewTestContext(dummy)
	opts := []string{"^\\dtest\\.$"}

	err := RegexpValidator(ctx, opts)

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "regexp.mustMatchPattern" {
		t.Fatalf("Expected must equal value error.")
	}
}

func TestThatRegexpValidatorFailsForStringNilValue(t *testing.T) {
	var dummy *string

	ctx := core.NewTestContext(dummy)
	opts := []string{"^\\dtest\\.$"}

	err := RegexpValidator(ctx, opts)

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "regexp.mustMatchPattern" {
		t.Fatalf("Expected must equal value error.")
	}
}

func TestThatRegexpValidatorFailsForUnsupportedValueType(t *testing.T) {
	type Dummy struct{}

	ctx := core.NewTestContext(&Dummy{})
	err := NumericValidator(ctx, []string{})

	if err.Error() != "type.unsupported" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
