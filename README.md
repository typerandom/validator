# Validator [![GoDoc](https://godoc.org/github.com/typerandom/validator?status.png)](http://godoc.org/github.com/typerandom/validator) [![Build Status](https://travis-ci.org/typerandom/validator.svg?branch=master)](https://travis-ci.org/typerandom/validator)

A powerful validation library for Go.

## Features

* Tag syntax that allows for typed parameters and multiple validation sets.
* Validation of deeply nested structures.
* Extensive list of [built-in validators](https://github.com/typerandom/validator/wiki/Validators).
* Localized error messages.
* Custom validators.

## Install

Just use go get.

    go get gopkg.in/typerandom/validator.v0
    
And then just import the package into your own code.

    import (
        "gopkg.in/typerandom/validator.v0"
    )

## Getting started

1. Add `validate` tags to the structure that you want to validate. See [Tagging](https://github.com/typerandom/validator/wiki/Tagging) and [Validators](https://github.com/typerandom/validator/wiki/Validators) for more details.
2. Call `errors := validator.Validate(objectWithValidateTags)`.
3. Call `errors.Any()` to check if there are any errors.
4. If there are errors, handle them. Or use `errors.PrintAll()` to print them to console (for debugging).
5. Questions? Check out the [wiki](https://github.com/typerandom/validator/wiki).

## Example


```go
package main

import (
	"gopkg.in/typerandom/validator.v0"
)

type User struct {
	Name  string `validate:"min(5),max(16)"`
	Email string `validate:"regexp(´^[a-z0-9-]*@[a-z0-9.]*\\.com$´)"`
	Age   int    `validate:"min(18),max(65)"`
}

func main() {
	user := &User{
		Name:  "Bob",
		Email: "bobby@tables",
		Age:   17,
	}

	if errs := validator.Validate(user); errs.Any() {
		errs.PrintAll()
		return
	}

	print("Hey " + user.Name + "!")
}
```

Running the example above would output:

    Name cannot be shorter than 5 characters.
    Email must match pattern '^[a-z0-9-]*@[a-z0-9.]*\.com$'.
    Age cannot be less than 18.

## Licensing

Validator is licensed under the MIT license. See LICENSE for the full license text.
