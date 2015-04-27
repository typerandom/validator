package main

import (
	"errors"
	"fmt"
	"strings"
)

var UnsupportedTypeError = errors.New("Validator not supported.")

func renderErrorMessage(field *reflectedField, tag *tag, message string) string {
	message = strings.Replace(message, "{field}", field.FullName(), 1)
	message = strings.Replace(message, "{validator}", tag.Name, 1)
	return message
}

func newValidatorError(field *reflectedField, tag *tag, err error) *ValidatorError {
	return &ValidatorError{
		Field:   field,
		Tag:     tag,
		message: renderErrorMessage(field, tag, err.Error()),
	}
}

type ValidatorError struct {
	Field   *reflectedField
	Tag     *tag
	message string
}

func (this *ValidatorError) String() string {
	return "{ error: " + this.message + "}"
}

func (this *ValidatorError) Error() string {
	return this.message
}

func NewErrors() *Errors {
	return &Errors{}
}

type Errors struct {
	Items []error
}

func (this *Errors) Add(err error) {
	this.Items = append(this.Items, err)
}

func (this *Errors) First() error {
	return this.Items[0]
}

func (this *Errors) Any() bool {
	return len(this.Items) > 0
}

func (this *Errors) PrintAll() {
	for _, err := range this.Items {
		fmt.Println(err)
	}
}
