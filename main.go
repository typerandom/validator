package main

type Person struct {
	FirstName string `validate:"empty,numeric,lowercase,min(2),max(64)"`
	//Age       int    `validate:"min(2),max(64)"`
}

func main() {
	var firstName string

	firstName = "1"

	person := &Person{FirstName: firstName}

	if errors := Validate(person); errors != nil {
		errors.PrintAll()
	}
}
