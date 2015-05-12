package parser

import (
	"errors"
	"strings"
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

type Method struct {
	Name       string
	Parameters []string
}

func (this *Method) String() string {
	params := strings.Join(this.Parameters, "', '")

	if len(params) == 0 {
		params = "(none)"
	} else {
		params = "'" + params + "'"
	}

	return "{ name: '" + this.Name + "', params: " + params + " }"
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
		case TOKEN_PARAM_FLOAT, TOKEN_PARAM_INTEGER, TOKEN_PARAM_STRING:
			method.Parameters = append(method.Parameters, token.value)
		case TOKEN_ERROR:
			return nil, errors.New(token.value)
		default:
			return nil, errors.New("Unable to parse. Unhandled token type.")
		}
	}

	methodGroups = append(methodGroups, methods)

	return methodGroups, nil
}
