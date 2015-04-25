package main

type Foo struct {
	Something string `validate:"min(3),max(6)"`
}

type Person struct {
	FirstName string  `validate:"min(15)"`
	Age       float32 `validate:"min(15)"`
	Moo       map[string][]Foo
}

func main() {
	person := &Person{
		FirstName: "Test Testersson",
		Age:       15,
		Moo: map[string][]Foo{
			"test": []Foo{Foo{Something: "13"}},
		},
	}

	if errors := Validate(person); errors.Any() {
		errors.PrintAll()
		return
	}

	print(person.FirstName)
}
