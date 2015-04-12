package main

func Register(name string, validator ValidatorStrategy) {
	if validators == nil {
		validators = make(map[string]ValidatorStrategy)
	}
	validators[name] = validator
}

func Validate(value interface{}) *Errors {
	var errors *Errors

	for _, field := range getTaggedFields(value, "validate") {
		for _, tag := range field.Tags {
			if validator, err := getValidator(tag.Name); err == nil {
				if err = validator(field.Value, tag.Options); err != nil {
					if errors == nil {
						errors = NewErrors()
					}
					errors.Add(NewValidatorError(field.Name, tag.Name, err.Error()))
				}
			} else {
				errors = NewErrors()
				errors.Add(NewValidatorError(field.Name, tag.Name, "Validator '"+tag.Name+"' used on field '"+field.Name+"' does not exist."))
				return errors
			}
		}
	}

	return errors
}
