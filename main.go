package main

type Person struct {
	FirstName *string `validate:"empty,min(2),max(64)"`
}

func main() {
	firstName := "Joe"
	person := &Person{FirstName: &firstName}
	Validate(person)
}

/*

A validation library

import (
	validator "github.com/typerandom/cocoon"
)

validator.RegisterValidator("not_empty", func(value string, params string) error {

})

type Foo struct {
	Id string `validate:"uuid"`
	Name string `validate:"regex(a-z0-9)"`
	Data string `validate:"func"`
	CreatedAt string `validate:"time"`
}

// Write your own validator method.
func (this *Foo) ValidateData(data string) error {
	if len(data) == 0 {

	}
	return nil
}

if errors := validator.Validate(foo); errors != nil {
	return errors.First()
}

*/
