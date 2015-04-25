package main

import (
	"fmt"
)

func NewValidatorError(fieldName string, tagName string, message string) *ValidatorError {
	return &ValidatorError{
		FieldName: fieldName,
		TagName:   tagName,
		Message:   message,
	}
}

type ValidatorError struct {
	FieldName string
	TagName   string
	Message   string
}

func (this *ValidatorError) String() string {
	return "{ error: " + this.Message + "}"
}

func (this *ValidatorError) Error() string {
	return this.Message
}

func NewErrors() *Errors {
	return &Errors{}
}

type Errors struct {
	Items []*ValidatorError
}

func (this *Errors) Add(err *ValidatorError) {
	this.Items = append(this.Items, err)
}

func (this *Errors) First() *ValidatorError {
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
