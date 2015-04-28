# Validator [![GoDoc](https://godoc.org/github.com/typerandom/validator?status.png)](http://godoc.org/github.com/typerandom/validator)

A simple, expressive and powerful validation library for Go.

## Features

* Tag structure fields with parameterizable validation rules.
* Validate deeply nested structures.
* Add custom validators.
* Load locale from JSON file.

## Getting Started

1. Import `github.com/typerandom/validator` into to your Go project.
2. Add validation tags to the structure that you want to validate. See section `Tagging` below.
3. Call `errors := validator.Validate(onYourObjectThatYouHaveGivenValidatorTags)`.
4. Check if the return value (errors) are empty (`errors.Any()`). If they are, then validation passed, if not, then show the errors to the user.

## Example

### example.go

	package main

	import (
		"github.com/typerandom/validator"
	)

	type Person struct {
		FirstName string `validate:"min(5),max(16)"`
		LastName  string `validate:"min(5),max(20)"`
		Email     string `validate:"regexp(^[a-z0-9-]*@[a-z0-9.]*\\\\.com$)"`
		Age       int    `validate:"min(18),max(65)"`
	}

	func main() {
		person := &Person{
			FirstName: "Bob",
			LastName:  "Tab",
			Email:     "bobby@tables",
			Age:       17,
		}

		if errors := validator.Validate(person); errors.Any() {
			errors.PrintAll()
			return
		}

		print("Hey " + person.FirstName + "!")
	}
	
Running the example above would output:

    FirstName cannot be shorter than 5 characters.
    LastName cannot be shorter than 5 characters.
    Email must match pattern '^[a-z0-9-]*@[a-z0-9.]*\.com$'.
    Age cannot be less than 18.

## Tagging

In order to specify how fields should be validated, fields must be tagged with the `validate` tag. The `validate` tag should contain validation rules in the format of `validator_name(params)` and should be separated by `,` for chaining of multiple rules. I.e. `some_validator,other_validator(abc)`.

### Examples

    Value int     `validate:"not_empty,max(10)"`
    Value string  `validate:"empty,regex(^[a-z_]*$),max(64)"`
    Value *string `validate:"not_empty,func"`

## Validators

### Minimum (`min`)

The lowest boundary of a type. I.e. the lowest a number or the length of a string can be.

#### Examples

    Value int               `validate:"min(5)"` // Value cannot be less than 5.
    Value string            `validate:"min(5)"` // Value cannot contain less than 5 characters.
    Value []string          `validate:"min(5)"` // Value cannot contain less than 5 items.
    Value map[string]string `validate:"min(5)"` // Value cannot contain less than 5 keys.
    
#### Supports

Strings, integers, floats, maps, arrays and slices.

### Maximum (`max`)

The highest boundary of a type. I.e. the highest a number or the length of a string can be.

#### Examples

    Value int               `validate:"max(5)"` // Value cannot be greater than 5.
    Value string            `validate:"max(5)"` // Value cannot be longer than 5 characters.
    Value []string          `validate:"max(5)"` // Value cannot contain more than 5 items.
    Value map[string]string `validate:"max(5)"` // Value cannot contain more than 5 keys.

### Not empty (`not_empty`)

Assert that a type is not empty. I.e. that a number or the length of a string is not 0, or that the value of a pointer is not `nil`.

#### Examples

    Value int               `validate:"not_empty"` // Value cannot be empty.
    Value *string           `validate:"not_empty"` // Value cannot be empty.
    Value []string          `validate:"not_empty"` // Value cannot be empty.
    Value map[string]string `validate:"not_empty"` // Value cannot be empty.
    
#### Supports

Pointers, strings, integers, floats, maps, arrays and slices.

### Not documented yet...

* empty
* equal
* contain
* lower_case
* upper_case
* numeric
* time
* regexp
* func

## Custom Validators

### Global

Register a global validator by calling `validator.Register(name string, validator ValidatorFn)`.

    validator.Register("validator_name", func (context *validator.Context, options []string) error {
        // ...
    })
    
### Local to struct

Use the `func` validator tag to assign a custom validator method on your structure. See `func` under the `Validators` section above for more information.

# License

MIT
