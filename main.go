package main

type Person struct {
	FirstName *string `validate:"empty,min(2),max(64)"`
}

func main() {
	firstName := "J"
	person := &Person{FirstName: &firstName}

	if errors := Validate(person); errors != nil {
		errors.PrintAll()
	}
}
