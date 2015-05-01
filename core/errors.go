package core

import (
	"fmt"
	"strings"
)

func NewValidatorError(field *ReflectedField, tag *Tag, err error) *Error {
	return &Error{
		Field:  field,
		Tag:    tag,
		Source: err,
	}
}

type Error struct {
	Field  *ReflectedField
	Tag    *Tag
	Source error
}

func (this *Error) String() string {
	return "{ error: " + this.Error() + "}"
}

func (this *Error) Error() string {
	message := strings.Replace(this.Source.Error(), "{field}", this.Field.FullName(), 1)
	message = strings.Replace(message, "{validator}", this.Tag.Name, 1)
	return message
}

type Errors struct {
	Items []error
}

func NewErrors() *Errors {
	return &Errors{}
}

func (this *Errors) Add(err error) {
	this.Items = append(this.Items, err)
}

func (this *Errors) AddMany(errs *Errors) {
	for _, err := range errs.Items {
		this.Items = append(this.Items, err)
	}
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
