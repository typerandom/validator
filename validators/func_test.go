package validators_test

import (
	"errors"
	"github.com/typerandom/validator/core"
	. "github.com/typerandom/validator/validators"
	"testing"
)

func newFuncTestContext(source interface{}, fieldName string) core.ValidatorContext {
	ctx := core.NewTestContext(nil)

	ctx.SetSource(source)

	ctx.SetField(&core.ReflectedField{
		Name:   fieldName,
		Parent: &core.ReflectedField{},
	})

	return ctx
}

func TestThatFuncValidatorFailsForInvalidOptions(t *testing.T) {
	ctx := core.NewTestContext(nil)

	err := FuncValidator(ctx, []string{"123", "123"})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "arguments.singleRequired" {
		t.Fatalf("Expected single argument required error.")
	}
}

func TestThatFuncValidatorFailsForMissingDefaultMethod(t *testing.T) {
	type Dummy struct {
		TestValue string
	}

	ctx := newFuncTestContext(&Dummy{}, "TestValue")
	err := FuncValidator(ctx, []string{})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "Validation method 'ValidateTestValue' on field '{field}' does not exist." {
		t.Fatalf("Expected invalid method error, got %s.", err)
	}
}

func TestThatFuncValidatorFailsForMissingExplicitMethod(t *testing.T) {
	type Dummy struct {
		TestValue string
	}

	ctx := newFuncTestContext(&Dummy{}, "TestValue")
	err := FuncValidator(ctx, []string{"TestSomeValue"})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "Validation method 'TestSomeValue' on field '{field}' does not exist." {
		t.Fatalf("Expected invalid method error, got %s.", err)
	}
}

type failingFuncDummy struct {
	TestValue string
}

func (f *failingFuncDummy) ValidateTestValue(context core.ValidatorContext) error {
	return errors.New("testvalue.error")
}

func (f *failingFuncDummy) TestSomeValue(context core.ValidatorContext) error {
	return errors.New("testvalue.error")
}

func (f *failingFuncDummy) TestInvalidInputParams(context core.ValidatorContext) {
}

func (f *failingFuncDummy) TestInvalidOutputParams(context core.ValidatorContext) (*string, error) {
	return nil, nil
}

func TestThatFuncValidatorFailsForExistingDefaultMethod(t *testing.T) {
	dummy := &failingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []string{})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "testvalue.error" {
		t.Fatalf("Expected explicit error result, got %s.", err)
	}
}

func TestThatFuncValidatorFailsForMissingExplicitMethodWithInvalidInputParams(t *testing.T) {
	dummy := &failingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []string{"TestInvalidInputParams"})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "Invalid return value(s) of validation method 'TestInvalidInputParams'. Return value must be of type 'error'." {
		t.Fatalf("Expected validation error result, got %s.", err)
	}
}

func TestThatFuncValidatorFailsForMissingExplicitMethodWithInvalidOutputParams(t *testing.T) {
	dummy := &failingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []string{"TestInvalidOutputParams"})

	if err == nil {
		t.Fatalf("Expected error, didn't get any.")
	}

	if err.Error() != "Invalid return value(s) of validation method 'TestInvalidOutputParams'. Return value must be of type 'error'." {
		t.Fatalf("Expected validation error result, got %s.", err)
	}
}

type passingFuncDummy struct {
	TestValue string
}

func (f *passingFuncDummy) ValidateTestValue(context core.ValidatorContext) error {
	return nil
}

func (f *passingFuncDummy) TestSomeValue(context core.ValidatorContext) error {
	return nil
}

func TestThatFuncValidatorSucceedsForExistingDefaultMethod(t *testing.T) {
	dummy := &passingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []string{})

	if err != nil {
		t.Fatalf("Didn't expect error, got %s.", err)
	}
}

func TestThatFuncValidatorSucceedsForExistingExplicitMethod(t *testing.T) {
	dummy := &passingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []string{"TestSomeValue"})

	if err != nil {
		t.Fatalf("Didn't expect error, got %s.", err)
	}
}
