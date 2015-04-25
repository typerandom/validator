package main

type Foo struct {
	Something string `validate:"min(3),max(6)"`
}

type Person struct {
	FirstName string         `validate:"min(15)"`
	Age       float32        `validate:"min(15)"`
	Foo       []Foo          `validate:"min(1),max(2)"`
	Moo       map[string]Foo `validate:"min(0),max(1)"`
}

func main() {
	person := &Person{
		FirstName: "Test Testersson",
		Age:       15,
		Foo: []Foo{
			Foo{Something: "HHehhe"},
			Foo{Something: "OHO"},
		},
	}

	if errors := Validate(person); errors.Any() {
		errors.PrintAll()
		return
	}

	print(person.FirstName)
}
