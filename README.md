# Validator [![GoDoc](https://godoc.org/github.com/typerandom/validator?status.png)](http://godoc.org/github.com/typerandom/validator) [![Build Status](https://travis-ci.org/typerandom/validator.svg?branch=master)](https://travis-ci.org/typerandom/validator)

A powerful validation library for Go.

## Features

* Tag syntax that allows for typed parameters and multiple validation sets.
* Validation of deeply nested structures.
* Custom validators.
* Localized error messages.

## Install

Just use go get.

    go get gopkg.in/typerandom/validator.v0
    
And then just import the package into your own code.

    import (
        "gopkg.in/typerandom/validator.v0"
    )

## Getting started

1. Add validation tags to the structure that you want to validate. See section `Tagging` below.
2. Call `errors := validator.Validate(yourObjectThatYouHaveGivenValidatorTags)`.
3. Call èrrors.Any()` to check if there are any errors.
4. If there are errors, handle them. Or use `errors.PrintAll()` if you're debugging.

## Example

### example.go

	package main

	import (
		"github.com/typerandom/validator"
	)

	type Person struct {
		FirstName string `validate:"min(5),max(16)"`
		LastName  string `validate:"min(5),max(20)"`
		Email     string `validate:"regexp(´^[a-z0-9-]*@[a-z0-9.]*\\.com$´)"`
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

In order to specify how fields should be validated, fields must be tagged with the `validate` tag. The `validate` tag should contain validation rules in the format of `validator_name(params)` and should be separated by `,` for chaining of multiple rules. I.e. `some_validator,other_validator(abc)`. Omitting `()` is the same as calling a method without parameters. I.e. `not_empty` is the same as `not_empty()`.

### Examples

    Value int     `validate:"not_empty,max(10)"`
    Value string  `validate:"empty,regex(^[a-z_]*$),max(64)"`
    Value *string `validate:"not_empty,func"`

# License

MIT
