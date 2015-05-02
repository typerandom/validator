package validators

import (
	"errors"
	"github.com/typerandom/validator/core"
	"testing"
)

func TestThatLowerCaseValidatorFailsForInvalidOptions(t *testing.T) {
	var dummy string

	ctx := core.NewTestContext(dummy)
	opts := []string{"123"}

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

	ctx := core.NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []string{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatLowerCaseValidatorSucceedsForNilString(t *testing.T) {
	var dummy *string

	ctx := core.NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []string{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatLowerCaseValidatorFailsForUpperCaseString(t *testing.T) {
	dummy := "ABC"

	ctx := core.NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []string{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "lowerCase.mustBeLowerCase" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatLowerCaseValidatorFailsForMixedCaseString(t *testing.T) {
	dummy := "aBc"

	ctx := core.NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []string{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "lowerCase.mustBeLowerCase" {
		t.Fatal(errors.New("Expected none arguments supported error."))
	}
}

func TestThatLowerCaseValidatorSucceedsForStringWithoutCase(t *testing.T) {
	dummy := "123"

	ctx := core.NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []string{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatLowerCaseValidatorSucceedsForLowerCaseString(t *testing.T) {
	dummy := "abc"

	ctx := core.NewTestContext(dummy)
	err := LowerCaseValidator(ctx, []string{})

	if err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func TestThatLowerCaseValidatorFailsForUnsupportedValueType(t *testing.T) {
	type Dummy struct{}

	ctx := core.NewTestContext(&Dummy{})
	err := LowerCaseValidator(ctx, []string{})

	if err.Error() != "type.unsupported" {
		t.Fatalf("Exepected unsupported type error, got %s.", err)
	}
}
