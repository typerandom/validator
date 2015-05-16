package main

import (
	"errors"
	"fmt"
	"github.com/typerandom/validator"
	"github.com/typerandom/validator/core"
)

// Validate using the 'func' validator and local validator method that is attached to the struct being validated.

type LocalUser struct {
	Name string `validate:"func"`
}

func (user *LocalUser) ValidateName(context core.ValidatorContext, args []interface{}) error {
	if user.Name != "test" {
		return errors.New("Name must be 'test'.")
	}
	return nil
}

func validateWithCustomLocalValidator() {
	fmt.Println("Validating using local 'func' validator...")

	user := &LocalUser{Name: "bob"}

	if errs := validator.Validate(user); errs.Any() {
		fmt.Printf("* Got error: %s\n", errs.First())
	}

	user.Name = "test"

	if errs := validator.Validate(user); !errs.Any() {
		fmt.Println("* User validation succeeded.")
	}

	fmt.Println()
}

// Validate using a global validator method that can be referenced in any validate tag (that uses the default validator).

type GlobalUser struct {
	Name string `validate:"globalTestValidator()"`
}

func validateWithCustomGlobalValidator() {
	fmt.Println("Validating using globally registered validator...")

	validator.Register("globalTestValidator", func(context core.ValidatorContext, args []interface{}) error {
		if len(args) != 0 {
			return context.NewError("arguments.noneSupported")
		}
		if stringValue, ok := context.Value().(string); ok {
			if stringValue != "test" {
				return errors.New("Name must be 'test'.")
			}
			return nil
		}
		return context.NewError("type.unsupported")
	})

	user := &GlobalUser{Name: "bob"}

	if errs := validator.Validate(user); errs.Any() {
		fmt.Printf("* Got error: %s\n", errs.First())
	}

	user.Name = "test"

	if errs := validator.Validate(user); !errs.Any() {
		fmt.Println("* User validation succeeded.")
	}
}

func main() {
	validateWithCustomLocalValidator()
	validateWithCustomGlobalValidator()
}
