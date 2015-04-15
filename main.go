package main

type Person struct {
	FirstName string `validate:"not_empty,min(1),max(6),lowercase"`
}

func main() {
	firstName := "Keeenleeeve"

	person := &Person{FirstName: firstName}

	if errors := Validate(person); errors != nil {
		errors.PrintAll()
		return
	}

	print(person.FirstName)
}
