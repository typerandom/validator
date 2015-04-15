package main

type Person struct {
	FirstName int `validate:"empty,min(2),max(64)"`
	//Age       int    `validate:"min(2),max(64)"`
}

func main() {
	var firstName int

	firstName = 1

	person := &Person{FirstName: firstName}

	if errors := Validate(person); errors != nil {
		errors.PrintAll()
	}
}
