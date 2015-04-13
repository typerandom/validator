package main

type Person struct {
	FirstName string `validate:"not_empty,min(2),max(64)"`
	Age       int    `validate:"not_empty,min(2),max(64)"`
}

func main() {
	var firstName string

	firstName = "b1"

	person := &Person{FirstName: firstName}

	/*registerStructFieldValidators(reflect.TypeOf(person), "name", IsEmpty)
	registerStructFieldValidators(reflect.TypeOf(otherPerson), "name", IsEmpty)
	print(len(typeValidatorRegistry))*/

	if errors := Validate(person); errors != nil {
		errors.PrintAll()
	}
}
