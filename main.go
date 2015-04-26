package main

import (
	"errors"
)

type Foo struct {
	Something string `validate:"min(3),max(6),func"`
}

func (this *Foo) ValidateSomething(context *ValidatorContext) error {
	if typedSomething, ok := context.Value.(string); ok {
		if typedSomething != "test" {
			return errors.New(context.Field.FullName() + " must be 'test'.")
		}
	}
	return nil
}

type Person struct {
	FirstName string  `validate:"min(15)"`
	Age       float32 `validate:"min(15),func"`
	Moo       map[string][]Foo
}

func (this *Person) ValidateAge(context *ValidatorContext) error {
	if typedAge, ok := context.Value.(float64); ok {
		if typedAge < 18 {
			return errors.New(context.Field.FullName() + " cannot be less than 18.")
		}
	}
	return nil
}

func main() {
	person := &Person{
		FirstName: "Test Testersson",
		Age:       18,
		Moo: map[string][]Foo{
			"test": []Foo{Foo{Something: "tes"}},
		},
	}

	if errors := Validate(person); errors.Any() {
		errors.PrintAll()
		return
	}

	print(person.FirstName)
}
