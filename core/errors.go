package core

import (
	"fmt"
	"github.com/typerandom/validator/core/parser"
	"strings"
)

type Error struct {
	field     *ReflectedField
	validator *parser.Method
	src       error
}

func NewError(field *ReflectedField, validator *parser.Method, err error) *Error {
	return &Error{
		field:     field,
		validator: validator,
		src:       err,
	}
}

func NewPlainError(err error) *Error {
	return &Error{
		src: err,
	}
}

func (this Error) IsFieldError() bool {
	return this.field != nil && this.validator != nil
}

func (this *Error) GetFieldName() string {
	if this.field == nil {
		return ""
	}
	return this.field.FullName()
}

func (this *Error) GetFieldDisplayName() string {
	if this.field == nil {
		return ""
	}
	return this.field.FullDisplayName()
}

func (this *Error) GetValidatorName() string {
	if this.validator == nil {
		return ""
	}
	return this.validator.Name
}

func (this *Error) String() string {
	return this.Error()
}

func (this *Error) Error() string {
	if this.IsFieldError() {
		message := strings.Replace(this.src.Error(), "{field}", this.GetFieldDisplayName(), 1)
		message = strings.Replace(message, "{validator}", this.GetValidatorName(), 1)
		return message
	} else {
		return this.src.Error()
	}
}

type ErrorList []*Error

func (this *ErrorList) AddPlain(err error) {
	this.Add(NewPlainError(err))
}

func (this *ErrorList) Add(err *Error) {
	*this = append(*this, err)
}

func (this *ErrorList) Clear() {
	*this = ErrorList{}
}

func (this *ErrorList) AddMany(errs ErrorList) {
	for _, err := range errs {
		this.Add(err)
	}
}

func (this ErrorList) WithField(fieldName string) ErrorList {
	var errs ErrorList

	for _, err := range this {
		if err.IsFieldError() && err.GetFieldName() == fieldName {
			errs.Add(err)
		}
	}

	return errs
}

func (this ErrorList) Any() bool {
	return len(this) > 0
}

func (this ErrorList) First() error {
	if this.Any() {
		return this[0]
	}
	return nil
}

func (this ErrorList) PrintAll() {
	for _, err := range this {
		fmt.Println(err)
	}
}
