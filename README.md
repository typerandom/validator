# Validator [![GoDoc](https://godoc.org/github.com/typerandom/validate?status.png)](http://godoc.org/github.com/typerandom/validate)

A validation library for Go.

------------------

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

The lowest that a number or the length of a string can be.

#### Examples

    Age int       `validate:"min(3)"`  // Age cannot be less than 3.
    Name string   `validate:"min(3)"`  // Name cannot contain less than 3 characters.
    Name []string `validate:"min(3)"`  // Name cannot contain less than 3 items.
    Name ]string  `validate:"min(3)"`  // Name cannot contain less than 3 keys.
    
#### Supports

Strings, integers, floats, maps, arrays and slices.

# License

MIT
