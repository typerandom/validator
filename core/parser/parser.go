package parser

import (
	"errors"
	"fmt"
	"strconv"
)

type Methods []*Method

func (this Methods) String() string {
	result := ""

	for _, method := range this {
		if result != "" {
			result += ", "
		}
		result += method.String()
	}

	return result
}

type Arguments []interface{}

func (args Arguments) String() string {
	if len(args) == 0 {
		return "(none)"
	}

	result := ""

	for _, val := range args {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf("'%v'", val)
	}

	return result
}

type Method struct {
	Name      string
	Arguments Arguments
}

func (this *Method) String() string {
	return "{ name: '" + this.Name + "', args: " + this.Arguments.String() + " }"
}

func Parse(text string) ([]Methods, error) {
	scanner := &scanner{
		value: text,
	}

	if len(scanner.value) > 0 {
		for lexer := lexGroup; lexer != nil; {
			lexer = lexer(scanner)
		}
	}

	var methodGroups []Methods
	var methods Methods
	var method *Method

	for _, token := range scanner.tokens {
		switch token.type_ {
		case TOKEN_GROUP:
			methodGroups = append(methodGroups, methods)
			methods = Methods{}
		case TOKEN_METHOD:
			method = &Method{
				Name: token.value,
			}
			methods = append(methods, method)
		case TOKEN_ARG_INTEGER, TOKEN_ARG_FLOAT:
			parsedValue, err := strconv.ParseFloat(token.value, 64)

			if err != nil {
				return nil, errors.New(fmt.Sprintf("Argument '%s' at position %d is not a valid number.", token.value, token.position))
			}

			method.Arguments = append(method.Arguments, parsedValue)
		case TOKEN_ARG_BOOLEAN:
			parsedValue, err := strconv.ParseBool(token.value)

			if err != nil {
				return nil, errors.New(fmt.Sprintf("Argument '%s' at position %d is not a valid boolean.", token.value, token.position))
			}

			method.Arguments = append(method.Arguments, parsedValue)
		case TOKEN_ARG_NIL:
			method.Arguments = append(method.Arguments, nil)
		case TOKEN_ARG_STRING:
			method.Arguments = append(method.Arguments, token.value)
		case TOKEN_ERROR:
			return nil, errors.New(token.value)
		default:
			return nil, errors.New("Unable to parse. Unhandled token type.")
		}
	}

	methodGroups = append(methodGroups, methods)

	return methodGroups, nil
}
