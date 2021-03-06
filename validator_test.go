package validator_test

import (
	. "github.com/typerandom/validator"
	"github.com/typerandom/validator/core"
	"testing"
)

func TestThatValidatorDefaultIsNotNil(t *testing.T) {
	if Default() == nil {
		t.Fatalf("Expected non nil value, got nil value.")
	}
}

func TestThatValidatorDefaultResolvesToSameInstance(t *testing.T) {
	validatorA := Default()
	validatorB := Default()

	if validatorA == nil || validatorB == nil {
		t.Fatalf("Expected non nil value, got nil value.")
	}

	if validatorA != validatorB {
		t.Fatalf("Expected resolved global validators to be same, got different.")
	}
}

func TestThatValidatorNewReturnsNewValidator(t *testing.T) {
	validatorA := New()
	validatorB := New()

	if validatorA == nil || validatorB == nil {
		t.Fatalf("Expected non nil value, got nil value.")
	}

	if validatorA == validatorB {
		t.Fatalf("Expected different validators, got same.")
	}
}

func TestThatValidatorCopyReturnsNewCopyOfValidator(t *testing.T) {
	validatorA := New()
	validatorB := validatorA.Copy()

	if validatorA == nil || validatorB == nil {
		t.Fatalf("Expected non nil value, got nil value.")
	}

	if validatorA == validatorB {
		t.Fatalf("Expected different validators, got same.")
	}

	if validatorA.Locale() == validatorB.Locale() {
		t.Fatalf("Expected different locales, got same.")
	}
}

func TestThatValidatorCanRegisterAndValidateCustomFunc(t *testing.T) {
	Default().Locale().Set("test.isNotTest", "test.isNotTest")

	Register("is_test", func(ctx core.ValidatorContext, args []interface{}) error {
		if val, ok := ctx.Value().(string); ok {
			if val == "test" {
				return nil
			}
			return ctx.NewError("test.isNotTest")
		}
		return ctx.NewError("type.unsupported")
	})

	type Dummy struct {
		Value string `validate:"is_test"`
	}

	errs := Validate(&Dummy{})

	if !errs.Any() {
		t.Fatalf("Expected error, didn't get any")
		return
	}

	if errs.First().Error() != "test.isNotTest" {
		t.Fatalf("Expected is not test error, got %s.", errs.First())
	}

	if errs = Validate(&Dummy{Value: "test"}); errs.Any() {
		t.Fatalf("Didn't expect error, got %s.", errs.First())
	}
}

func TestThatValidatorCanValidateStructValue(t *testing.T) {
	type Dummy struct {
		Value *string `validate:"nil|equal(test)|equal(other_test)"`
	}

	if errs := Validate(&Dummy{Value: nil}); errs.Any() {
		t.Fatalf("Didn't expect error, got %s.", errs.First())
	}

	dummyValue := "test"

	if errs := Validate(&Dummy{Value: &dummyValue}); errs.Any() {
		t.Fatalf("Didn't expect error, got %s.", errs.First())
	}

	dummyValue = "other_test"

	if errs := Validate(&Dummy{Value: &dummyValue}); errs.Any() {
		t.Fatalf("Didn't expect error, got %s.", errs.First())
	}

	dummyValue = "invalid_value"

	if errs := Validate(&Dummy{Value: &dummyValue}); !errs.Any() {
		t.Fatalf("Expected error, didn't get any")
	}
}

func TestThatValidatorDoesntValidateNilStructValue(t *testing.T) {
	type DummyStruct struct {
		Value string `validate:"not_empty"`
	}

	type Dummy struct {
		NonNilStruct *DummyStruct
		NilStruct    *DummyStruct
	}

	dummyValue := &Dummy{
		NonNilStruct: &DummyStruct{},
	}

	errs := Validate(dummyValue)

	if !errs.Any() {
		t.Fatal("Expected errors, but didn't get any.")
	}

	if errs.Length() != 1 {
		t.Fatalf("Expected 1 error, but got %d.", errs.Length())
	}

	firstError := errs.First()

	if firstError.String() != "NonNilStruct.Value cannot be empty." {
		t.Fatalf("Expected error to be 'NonNilStruct.Value cannot be empty.' but it was '%s'.", firstError.String())
	}
}
