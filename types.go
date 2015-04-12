package main

type ValidationError interface {
	FieldName() string
	ErrorType() string
	Error() string
}

func NewErrors() *Errors {
	return &Errors{}
}

type Errors struct {
	Items []ValidationError
}

func (this *Errors) Add(err ValidationError) {
	return
}

func (this *Errors) First() ValidationError {
	return this.Items[0]
}
