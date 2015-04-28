package core

import (
	"fmt"
)

type Errors struct {
	Items []error
}

func NewErrors() *Errors {
	return &Errors{}
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
