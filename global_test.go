package validator

import (
	"testing"
)

func TestThatGlobalValidatorIsNotNil(t *testing.T) {
	if getGlobalValidator() == nil {
		t.Fatalf("Expected non nil value, got nil value.")
	}
}

func TestThatGlobalValidatorResolvesToSameInstance(t *testing.T) {
	validatorA := getGlobalValidator()
	validatorB := getGlobalValidator()

	if validatorA == nil || validatorB == nil {
		t.Fatalf("Expected non nil value, got nil value.")
	}

	if validatorA != validatorB {
		t.Fatalf("Expected resolved global validators to be same, got different.")
	}
}
