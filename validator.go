package main

func Validate(value interface{}) *Errors {
	//errors := NewErrors()

	for _, field := range getTaggedFields(value, "validate") {
		print(field.Name)
		print("\n")
		if stringValue, ok := field.Value.(*string); ok {
			print(*stringValue)
		} else {
			print("Unknown value")
		}
		print("\n")
		print(field.Tag)
		print("\n-----------------\n")
	}

	return nil
}
