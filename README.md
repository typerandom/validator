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

1. Add validation tags to the structure that you want to validate. See [Tagging](https://github.com/typerandom/validator/wiki/Tagging) for more details.
2. Call `errors := validator.Validate(objectThatYouHaveGivenValidatorTags)`.
3. Call `errors.Any()` to check if there are any errors.
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

## Licensing

Validator is licensed under the MIT license. See LICENSE for the full license text.
