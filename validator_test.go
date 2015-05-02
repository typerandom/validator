package validator

import (
	"github.com/typerandom/validator/core"
	"testing"
)

func TestThatValidatorDefaultIsNotNil(t *testing.T) {
	if Default() == nil {
		t.Fatalf("Expected non nil value, got nil value.")
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

func TestThatValidatorRegisterCanRegisterAndValidateWithCustomFunc(t *testing.T) {
	Default().locale.Set("test.isNotTest", "test.isNotTest")

	Register("is_test", func(ctx core.ValidatorContext, options []string) error {
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
