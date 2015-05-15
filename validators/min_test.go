package validators_test

import (
	"errors"
	. "github.com/typerandom/validator/testing"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func TestThatMinValidatorFailsForInvalidOptions(t *testing.T) {
	dummy := 100

	ctx := NewTestContext(dummy)
	err := MinValidator(ctx, []interface{}{})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatal(errors.New("Expected single argument required error."))
	}

	err = MinValidator(ctx, []interface{}{"abc"})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.invalidType" {
		t.Fatal(errors.New("Expected invalid arguments error."))
	}

	err = MinValidator(ctx, []interface{}{"123", "123"})

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatal(errors.New("Expected single argument required error."))
	}
}

func testThatMinValidatorSucceedsForValueOverLimit(t *testing.T, limit float64, dummy interface{}) {
	ctx := NewTestContext(dummy)
	opts := []interface{}{limit}

	if err := MinValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func testThatMinValidatorSucceedsForValueOnLimit(t *testing.T, limit float64, dummy interface{}) {
	ctx := NewTestContext(dummy)
	opts := []interface{}{limit}

	if err := MinValidator(ctx, opts); err != nil {
		t.Fatalf("Didn't expect error, but got one (%s).", err)
	}
}

func testThatMinValidatorFailsForValueUnderLimit(t *testing.T, limit float64, dummy interface{}, expectedErr string) {
	ctx := NewTestContext(dummy)
	opts := []interface{}{limit}

	err := MinValidator(ctx, opts)

	if err == nil {
		t.Fatal(errors.New("Expected error, didn't get any."))
	}

	if err.Error() != expectedErr {
		t.Fatal(errors.New("Expected cannot be less than error."))
	}
}

func TestThatMinValidatorSucceedsForIntValueOverLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOverLimit(t, 5, 6)
}

func TestThatMinValidatorSucceedsForIntValueOnLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOnLimit(t, 5, 5)
}

func TestThatMinValidatorFailsForIntValueUnderLimit(t *testing.T) {
	testThatMinValidatorFailsForValueUnderLimit(t, 5, 4, "min.cannotBeLessThan")
}

func TestThatMinValidatorSucceedsForFloatValueOverLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOverLimit(t, 5.5, 5.6)
}

func TestThatMinValidatorSucceedsForFloatValueOnLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOnLimit(t, 5.5, 5.5)
}

func TestThatMinValidatorFailsForFloatValueUnderLimit(t *testing.T) {
	testThatMinValidatorFailsForValueUnderLimit(t, 5.5, 5.4, "min.cannotBeLessThan")
}

func TestThatMinValidatorSucceedsForStringValueOverLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOverLimit(t, 5, "123456")
}

func TestThatMinValidatorSucceedsForStringValueOnLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOnLimit(t, 5, "12345")
}

func TestThatMinValidatorFailsForStringValueUnderLimit(t *testing.T) {
	testThatMinValidatorFailsForValueUnderLimit(t, 5, "1234", "min.cannotBeShorterThan")
}

func TestThatMinValidatorSucceedsForSliceLengthOverLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOverLimit(t, 5, []string{"1", "2", "3", "4", "5", "6"})
}

func TestThatMinValidatorSucceedsForSliceLengthOnLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOnLimit(t, 5, []string{"1", "2", "3", "4", "5"})
}

func TestThatMinValidatorFailsForSliceLengthUnderLimit(t *testing.T) {
	testThatMinValidatorFailsForValueUnderLimit(t, 5, []string{"1", "2", "3", "4"}, "min.cannotContainLessItemsThan")
}

func TestThatMinValidatorSucceedsForMapLengthOverLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOverLimit(t, 5, map[string]string{"1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6"})
}

func TestThatMinValidatorSucceedsForMapLengthOnLimit(t *testing.T) {
	testThatMinValidatorSucceedsForValueOnLimit(t, 5, map[string]string{"1": "1", "2": "2", "3": "3", "4": "4", "5": "5"})
}

func TestThatMinValidatorFailsForMapLengthUnderLimit(t *testing.T) {
	testThatMinValidatorFailsForValueUnderLimit(t, 5, map[string]string{"1": "1", "2": "2", "3": "3", "4": "4"}, "min.cannotContainLessKeysThan")
}

func TestThatMinValidatorFailsForUnsupportedType(t *testing.T) {
	type Dummy struct{}
	testThatMinValidatorFailsForValueUnderLimit(t, 5, &Dummy{}, "type.unsupported")
}
