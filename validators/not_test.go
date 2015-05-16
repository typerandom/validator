package validators_test

import (
	"errors"
	"github.com/typerandom/validator/core"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatNotValidatorFailsForInvalidOptions(t *testing.T) {
	dummy := float64(100)

	ctx := core.NewTestContext(dummy)
	err := NotValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatal(errors.New("Expected single argument required error."))
	}

	err = NotValidator(ctx, []interface{}{float64(123)})

	if err != nil {
		t.Fatalf("Didn't expect error, but got '%s'.", err)
	}

	err = NotValidator(ctx, []interface{}{"123", "123"})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatal(errors.New("Expected single argument required error."))
	}
}

func testThatNotValidatorSucceedsWhenValueIsNotArgumentNotValue(t *testing.T, value interface{}, notValue interface{}) {
	ctx := core.NewTestContext(value)
	err := NotValidator(ctx, []interface{}{notValue})

	if err != nil {
		t.Fatalf("Didn't expect any error, but got '%s'.", err)
	}
}

func testThatNotValidatorFailsWhenValueIsArgumentNotValue(t *testing.T, value interface{}, notValue interface{}, expectedError string) {
	ctx := core.NewTestContext(value)
	err := NotValidator(ctx, []interface{}{notValue})

	if err == nil {
		t.Fatal("Expected error, but didn't get any.")
	}

	if err.Error() != expectedError {
		t.Fatalf("Expected '%s' error, but got '%s'.", expectedError, err)
	}
}

func TestThatNotValidatorFailsWhenNilValueIsNotArgumentNotValue(t *testing.T) {
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, nil, nil, "not.cannotBeValue")
}

func TestThatNotValidatorSucceedsWhenStringValueIsNotArgumentValue(t *testing.T) {
	testThatNotValidatorSucceedsWhenValueIsNotArgumentNotValue(t, "abc", nil)
	testThatNotValidatorSucceedsWhenValueIsNotArgumentNotValue(t, "abc", 123)
	testThatNotValidatorSucceedsWhenValueIsNotArgumentNotValue(t, "abc", "def")
}

func TestThatNotValidatorFailsWhenStringValueIsNotArgumentNotValue(t *testing.T) {
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, "abc", "abc", "not.cannotBeValue")
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, "123", 123, "not.cannotBeValue")
}

func TestThatNotValidatorSucceedsWhenIntValueIsNotIntArgumentValue(t *testing.T) {
	testThatNotValidatorSucceedsWhenValueIsNotArgumentNotValue(t, int64(123), float64(456))
	testThatNotValidatorSucceedsWhenValueIsNotArgumentNotValue(t, int64(456), float64(123))
}

func TestThatNotValidatorFailsWhenIntValueIsNotArgumentNotValue(t *testing.T) {
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, int64(123), float64(123), "not.cannotBeValue")
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, int64(456), float64(456), "not.cannotBeValue")
}

func TestThatNotValidatorSucceedsWhenFloatValueIsNotIntArgumentValue(t *testing.T) {
	testThatNotValidatorSucceedsWhenValueIsNotArgumentNotValue(t, float64(123), float64(456))
	testThatNotValidatorSucceedsWhenValueIsNotArgumentNotValue(t, float64(456), float64(123))
}

func TestThatNotValidatorFailsWhenFloatValueIsNotArgumentNotValue(t *testing.T) {
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, float64(123), float64(123), "not.cannotBeValue")
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, float64(456), float64(456), "not.cannotBeValue")
}

func TestThatNotValidatorFailsWhenValueIsUnsupported(t *testing.T) {
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, float64(123), "abc", "type.unsupported")
	testThatNotValidatorFailsWhenValueIsArgumentNotValue(t, int64(123), "abc", "type.unsupported")
}
