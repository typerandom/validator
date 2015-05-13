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

func TestThatFuncValidatorFailsForMissingDefaultMethod(t *testing.T) {
	type Dummy struct {
		TestValue string
	}

	ctx := newFuncTestContext(&Dummy{}, "TestValue")
	err := FuncValidator(ctx, []interface{}{})

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
	err := FuncValidator(ctx, []interface{}{"TestSomeValue"})

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

func (f *failingFuncDummy) ValidateTestValue(context core.ValidatorContext, args []interface{}) error {
	return errors.New("testvalue.error")
}

func (f *failingFuncDummy) TestSomeValue(context core.ValidatorContext, args []interface{}) error {
	return errors.New("testvalue.error")
}

func (f *failingFuncDummy) TestInvalidInputParams(context core.ValidatorContext, args []interface{}) {
}

func (f *failingFuncDummy) TestInvalidOutputParams(context core.ValidatorContext, args []interface{}) (*string, error) {
	return nil, nil
}

func TestThatFuncValidatorFailsForExistingDefaultMethod(t *testing.T) {
	dummy := &failingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []interface{}{})

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
	err := FuncValidator(ctx, []interface{}{"TestInvalidInputParams"})

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
	err := FuncValidator(ctx, []interface{}{"TestInvalidOutputParams"})

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

func (f *passingFuncDummy) ValidateTestValue(context core.ValidatorContext, args []interface{}) error {
	return nil
}

func (f *passingFuncDummy) TestSomeValue(context core.ValidatorContext, args []interface{}) error {
	return nil
}

func (f *passingFuncDummy) TestSomeValueWithArgs(context core.ValidatorContext, args []interface{}) error {
	if len(args) == 0 {
		panic("Expected arguments. Didn't get any!")
	}

	if t, ok := args[0].(*testing.T); ok {
		if val1, ok := args[1].(float64); ok {
			if val1 != 141 {
				t.Fatalf("Expected argument 1 to have value '%d', but got '%v'.", 141, val1)
			}
		} else {
			t.Fatal("Expected argument 1 to be float, but got something else...")
		}

		if val2, ok := args[2].(bool); ok {
			if val2 != true {
				t.Fatalf("Expected argument 2 to have value '%v', but got '%v'.", true, val2)
			}
		} else {
			t.Fatal("Expected argument 2 to be boolean, but got something else...")
		}
	} else {
		panic("Expected test argument. Got something else!")
	}

	return nil
}

func TestThatFuncValidatorSucceedsForExistingDefaultMethod(t *testing.T) {
	dummy := &passingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []interface{}{})

	if err != nil {
		t.Fatalf("Didn't expect error, got %s.", err)
	}
}

func TestThatFuncValidatorSucceedsForExistingExplicitMethod(t *testing.T) {
	dummy := &passingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []interface{}{"TestSomeValue"})

	if err != nil {
		t.Fatalf("Didn't expect error, got %s.", err)
	}
}

func TestThatFuncValidatorSucceedsForExistingExplicitWithArgumentsMethod(t *testing.T) {
	dummy := &passingFuncDummy{}

	ctx := newFuncTestContext(dummy, "TestValue")
	err := FuncValidator(ctx, []interface{}{"TestSomeValueWithArgs", t, float64(141), true})

	if err != nil {
		t.Fatalf("Didn't expect error, got %s.", err)
	}
}
