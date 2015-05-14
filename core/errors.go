package core

import (
	"fmt"
	"github.com/typerandom/validator/core/parser"
	"strings"
)

func NewError(field *ReflectedField, tag *parser.Method, err error) *Error {
	return &Error{
		Field:  field,
		Tag:    tag,
		Source: err,
	}
}

type Error struct {
	Field  *ReflectedField
	Tag    *parser.Method
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

type ErrorList []error

func (this *ErrorList) Add(err error) {
	*this = append(*this, err)
}

func (this *ErrorList) AddMany(errs ErrorList) {
	for _, err := range errs {
		this.Add(err)
	}
}

func (this ErrorList) First() error {
	return this[0]
}

func (this ErrorList) Any() bool {
	return len(this) > 0
}

func (this ErrorList) PrintAll() {
	for _, err := range this {
		fmt.Println(err)
	}
}
