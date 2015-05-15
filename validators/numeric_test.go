package validators_test

import (
	"errors"
	"github.com/typerandom/validator/core"
	. "github.com/typerandom/validator/testing"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatNumericValidatorFailsForInvalidOptions(t *testing.T) {
	var dummy *string

	ctx := NewTestContext(dummy)
	opts := []interface{}{"123"}

	err := NumericValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.noneSupported" {
		t.Fatalf("Expected none arguments supported error.")
	}
}

func testThatNumericValidatorSucceedsForValue(t *testing.T, value interface{}) core.ValidatorContext {
	ctx := NewTestContext(value)
	opts := []interface{}{}

	if err := NumericValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}

	return ctx
}

func TestThatNumericValidatorSucceedsForIntValue(t *testing.T) {
	ctx := testThatNumericValidatorSucceedsForValue(t, 5)

	if _, ok := ctx.Value().(int64); !ok {
		t.Fatalf("Expected int64, recieved other type.")
	}
}

func TestThatNumericValidatorSucceedsForFloatValue(t *testing.T) {
	ctx := testThatNumericValidatorSucceedsForValue(t, 5.5)

	if _, ok := ctx.Value().(float64); !ok {
		t.Fatalf("Expected float64, recieved other type.")
	}
}

func TestThatNumericValidatorSucceedsForIntStringValue(t *testing.T) {
	ctx := testThatNumericValidatorSucceedsForValue(t, "5")

	if _, ok := ctx.Value().(float64); !ok {
		t.Fatalf("Expected float64, recieved other type.")
	}
}

func TestThatNumericValidatorSucceedsForFloatStringValue(t *testing.T) {
	ctx := testThatNumericValidatorSucceedsForValue(t, "5.5")

	if _, ok := ctx.Value().(float64); !ok {
		t.Fatalf("Expected float64, recieved other type.")
	}
}

func TestThatNumericValidatorFailsForUnsupportedValueType(t *testing.T) {
	type Dummy struct{}

	ctx := NewTestContext(&Dummy{})
	err := NumericValidator(ctx, []interface{}{})

	if err.Error() != "type.unsupported" {
		t.Fatalf("Expected unsupported type error, got %s.", err)
	}
}
