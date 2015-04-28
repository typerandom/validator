# Validator [![GoDoc](https://godoc.org/github.com/typerandom/validate?status.png)](http://godoc.org/github.com/typerandom/validate)

A simple, expressive and powerful validation library for Go.

## Features

* Tag structure fields with parameterizable validation rules.
* Validate deeply nested structures.
* Add custom validators.
* Load locale from JSON file.

## Getting Started

## Example

	package main

	import (
		"github.com/typerandom/validator"
	)

	type Person struct {
		FirstName string `validate:"min(2),max(16)"`
		LastName  string `validate:"min(2),max(20)"`
		Email     string `validate:"regexp(^[a-z0-9-]*@[a-z0-9.]*\\\\.com$)"`
		Age       int    `validate:"min(18),max(65)"`
	}

	func main() {
		person := &Person{
			FirstName: "Bobby",
			LastName:  "Tables",
			Email:     "bob@tables.com",
			Age:       19,
		}

		if errors := validator.Validate(person); errors.Any() {
			errors.PrintAll()
			return
		}

		print("Hey " + person.FirstName + "!")
	}


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

# License

MIT
