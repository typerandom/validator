package bench

import (
	"errors"
	asaskevichGoValidator "github.com/asaskevich/govalidator"
	validator "github.com/typerandom/validator"
	goValidator "gopkg.in/validator.v2"
	"strconv"
	"testing"
)

var nativeError error

func BenchmarkNativeMin(b *testing.B) {
	type Foo struct {
		StringValue string
		IntValue    int
	}

	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}
	invalidFoo := &Foo{StringValue: "Fo", IntValue: 3}

	foos := []*Foo{validFoo, invalidFoo}

	for i := 0; i < b.N; i++ {
		for j := 0; j < 2; j++ {
			nativeError = nil
			foo := foos[j%2]

			if len(foo.StringValue) < 5 || len(foo.StringValue) > 10 {
				nativeError = errors.New("StringValue must be between 5 and 10 characters.")
			}
			if foo.IntValue < 5 || foo.IntValue > 10 {
				nativeError = errors.New("IntValue must be between 5 and 10.")
			}
		}
	}
}

func BenchmarkValidatorMin(b *testing.B) {
	type Foo struct {
		StringValue string `validate:"min(5),max(10)"`
		IntValue    int    `validate:"min(5),max(10)"`
	}

	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}
	invalidFoo := &Foo{StringValue: "Fo", IntValue: 3}

	for i := 0; i < b.N; i++ {
		validator.Validate(validFoo)
		validator.Validate(invalidFoo)
	}
}

func BenchmarkCompetitorGoValidatorMin(b *testing.B) {
	type Foo struct {
		StringValue string `validate:"min=5,max=10"`
		IntValue    int    `validate:"min=5,max=10"`
	}

	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}
	invalidFoo := &Foo{StringValue: "Fo", IntValue: 3}

	for i := 0; i < b.N; i++ {
		goValidator.Validate(validFoo)
		goValidator.Validate(invalidFoo)
	}
}

func BenchmarkCompetitorAsaskevichGoValidatorMin(b *testing.B) {
	type Foo struct {
		StringValue string `validate:"min_str"`
		IntValue    int    `validate:"min_int"`
	}

	asaskevichGoValidator.TagMap["min_str"] = asaskevichGoValidator.Validator(func(value string) bool {
		if len(value) < 5 || len(value) > 10 {
			return false
		}
		return true
	})

	asaskevichGoValidator.TagMap["min_int"] = asaskevichGoValidator.Validator(func(value string) bool {
		intValue, _ := strconv.ParseInt(value, 10, 32)

		if intValue < 5 || intValue > 10 {
			return false
		}

		return true
	})

	validFoo := &Foo{StringValue: "Foobar", IntValue: 7}
	invalidFoo := &Foo{StringValue: "Fo", IntValue: 3}

	for i := 0; i < b.N; i++ {
		asaskevichGoValidator.ValidateStruct(validFoo)
		asaskevichGoValidator.ValidateStruct(invalidFoo)
	}
}
