package main

type Person struct {
	FirstName *string `validate:"min(15)"`
}

func main() {
	firstName := "Keeenleeeve"

	person := &Person{FirstName: &firstName}

	if errors := Validate(person); errors != nil {
		errors.PrintAll()
		return
	}

	print(person.FirstName)
}
