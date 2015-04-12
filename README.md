cocoon
------------------

Cocoon is a validation library for Go.

**NOTE: WORK IN PROGRESS! SOME OF THE FEATURES LISTED BELOW MIGHT NOT WORK/BE SUPPORTED YET!**

# Features

* Large set of predefined validation methods.
* Validate structures using field tags. Recursive, so that structures of structures can even be validated.
* Customizable. Register your own validation methods. With the validation tag `func` you can provide your own valdation method on the struct (`Validate{FieldName}`).
* Expressional. Using the `validate` tag you can specify how a field on a struct should be validated.

# API

## Methods

### func Validate(value interface{}) \*Errors

Validate fields on a structure.

    type Something struct {
        Id    *string `validate:"empty,uuid"`
        Name  string  `validate:"not_empty,alpha,min(2),max(64)"
        Email string  `validate:"not_empty,email,func(ValidateUniqueEmail)"`
    }
    
    func (this *Something) ValidateUniqueEmail(email interface{}) error {
        if emailString, ok := email.(string); ok {
            ok, err := globalDatastore.IsEmailUnique(emailString)
            
            if err != nil {
                return err
            }
            
            if !ok {
                return NewValidationError(Email '"+emailString+"' is not unique.")
            }
        }
        return nil
    }
    
    func main() {
        something := &Something{
            Id:    "123",
            Name:  "bob",
            Email: "bob@bob.com",
        }
    
        if errors := cocoon.Validate(something); errors != nil {
            errors.PrintAll() // Field 'Id' is not a valid UUID.
            return
        }
    }
    
### func Register(name string, validator ValidatorFunc)

Register a validation function that can be referenced in struct tags.

    type Something struct {
        LaunchAt *string `validate:"date.iso8601"`
    }
    
    cocoon.Register("date.iso8601", func(value interface{}, options []string) error {
        // Validate that value is a string with an ISO 8601 format
    })
    
## struct Errors

### func PrintAll()

Outputs all of the errors to the console.

### func First() ValidationError

Get the first error in the list.

# License

MIT