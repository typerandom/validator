package main

type Person struct {
	FirstName string  `validate:"min(15)"`
	Age       float32 `validate:"min(15)"`
}

func main() {
	person := &Person{FirstName: "Test Testersson", Age: 15}

	if errors := Validate(person); errors != nil {
		errors.PrintAll()
		return
	}

	print(person.FirstName)
}
