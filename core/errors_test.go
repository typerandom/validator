package core

import (
	"errors"
	"github.com/typerandom/validator/core/parser"
	"testing"
)

func TestThatFieldErrorStringFormatIsValid(t *testing.T) {
	field := &ReflectedField{Name: "myField", DisplayName: "myField"}
	validator := &parser.Method{Name: "myValidator"}

	err := NewError(field, validator, errors.New("Error: {field} on {validator}."))

	if expectedErr := "Error: myField on myValidator."; err.String() != expectedErr {
		t.Fatalf("Expected '%s', got '%s'.", expectedErr, err)
	}
}

func TestThatFieldErrorIsFieldError(t *testing.T) {
	field := &ReflectedField{}
	validator := &parser.Method{}

	err := NewError(field, validator, errors.New(""))

	if !err.IsFieldError() {
		t.Fatal("Expected error to be a field error, but it wasn't.")
	}
}

func TestThatPlainErrorStringFormatIsValid(t *testing.T) {
	expectedErr := "Error: just a plain error."

	err := NewPlainError(errors.New(expectedErr))

	if err.String() != expectedErr {
		t.Fatalf("Expected '%s', got '%s'.", expectedErr, err)
	}
}

func TestThatPlainErrorIsPlainError(t *testing.T) {
	err := NewPlainError(errors.New(""))

	if err.IsFieldError() {
		t.Fatal("Expected error to not be a field error, but it was.")
	}
}

func TestThatErrorListAnyReturnsFalseWhenThereAreNoErrors(t *testing.T) {
	var errs ErrorList

	if errs.Any() {
		t.Fatal("Didn't expect any errors, but got that there were ones.")
	}
}

func TestThatErrorListAnyReturnsTrueWhenThereArePlainErrors(t *testing.T) {
	var errs ErrorList

	err := NewPlainError(nil)
	errs.Add(err)

	if !errs.Any() {
		t.Fatal("Expected errors, but got that there wasn't any.")
	}

	if errs.First() != err {
		t.Fatal("Expected error added to be first error, but it wasn't.")
	}
}

func TestThatErrorListAnyReturnsTrueWhenThereAreValidatorErrors(t *testing.T) {
	var errs ErrorList

	err := NewError(nil, nil, nil)
	errs.Add(err)

	if !errs.Any() {
		t.Fatal("Expected errors, but got that there wasn't any.")
	}

	if errs.First() != err {
		t.Fatal("Expected error added to be first error, but it wasn't.")
	}
}

func TestThatErrorListAnyReturnsTrueWhenThereAreManyErrors(t *testing.T) {
	var errs ErrorList

	validatorErr := NewError(nil, nil, nil)
	plainErr := NewPlainError(nil)

	errs.AddMany(ErrorList{
		validatorErr,
		plainErr,
	})

	if !errs.Any() {
		t.Fatal("Expected errors, but got that there wasn't any.")
	}

	if len(errs) != 2 {
		t.Fatal("Expected there to be 2 errors, but there wasn't.")
	}

	if errs[0] != validatorErr {
		t.Fatal("Expected first error to be validator error, but it wasn't.")
	}

	if errs[1] != plainErr {
		t.Fatal("Expected first error to be plain error, but it wasn't.")
	}
}

func TestThatWhenFilledErrorListIsClearedThereAreNoErrorsLeft(t *testing.T) {
	var errs ErrorList

	errs.Add(NewError(nil, nil, nil))

	if !errs.Any() {
		t.Fatal("Expected errors, but got that there wasn't any.")
	}

	errs.Clear()

	if errs.Any() {
		t.Fatalf("Didn't expect any errors, but there was %d.", len(errs))
	}
}

func TestThatErrorListForFieldOnlyReturnsTheSpecifiedField(t *testing.T) {
	var errs ErrorList

	parentField := &ReflectedField{Name: "User", DisplayName: "User"}

	errs.Add(NewPlainError(errors.New("Ooops.")))

	userFieldFirstNameError := NewError(
		&ReflectedField{Parent: parentField, Name: "FirstName", DisplayName: "FirstName"},
		&parser.Method{Name: "not_empty"},
		errors.New("'{field}' cannot be empty."),
	)

	errs.Add(userFieldFirstNameError)

	userFieldLastNameError := NewError(
		&ReflectedField{Parent: parentField, Name: "LastName", DisplayName: "LastName"},
		&parser.Method{Name: "not_empty"},
		errors.New("'{field}' cannot be empty."),
	)

	errs.Add(userFieldLastNameError)

	userFieldFirstNameErrors := errs.WithField("User.FirstName")

	if !userFieldFirstNameErrors.Any() {
		t.Fatal("Expected user field error, but didn't get any.")
	}

	if userFieldFirstNameErrors.First() != userFieldFirstNameError {
		t.Fatal("Expected user field error to be same as we created. But got something else.")
	}

	if len(userFieldFirstNameErrors) != 1 {
		t.Fatal("Expected one error, but got %d.", len(userFieldFirstNameErrors))
	}
}
