package main

import (
	"fmt"
	"github.com/typerandom/validator"
)

func validateString() {
	fmt.Println("Validating string value with min...")

	type Foo struct {
		Value string `validate:"min(4)"`
	}

	foo := &Foo{Value: ""}

	if errs := validator.Validate(foo); errs.Any() {
		fmt.Printf("* Got error: %s\n", errs.First())
	}

	foo.Value = "test"

	if errs := validator.Validate(foo); !errs.Any() {
		fmt.Println("* Foo validation succeeded.")
	}

	fmt.Println()
}

func validateInt() {
	fmt.Println("Validating int value with min...")

	type Foo struct {
		Value int `validate:"min(4)"`
	}

	foo := &Foo{Value: 0}

	if errs := validator.Validate(foo); errs.Any() {
		fmt.Printf("* Got error: %s\n", errs.First())
	}

	foo.Value = 4

	if errs := validator.Validate(foo); !errs.Any() {
		fmt.Println("* Foo validation succeeded.")
	}

	fmt.Println()
}

func validateArray() {
	fmt.Println("Validating array value with min...")

	type Foo struct {
		Value []int `validate:"min(4)"`
	}

	foo := &Foo{Value: []int{}}

	if errs := validator.Validate(foo); errs.Any() {
		fmt.Printf("* Got error: %s\n", errs.First())
	}

	foo.Value = []int{1, 2, 3, 4}

	if errs := validator.Validate(foo); !errs.Any() {
		fmt.Println("* Foo validation succeeded.")
	}
}

func main() {
	validateString()
	validateInt()
	validateArray()
}
