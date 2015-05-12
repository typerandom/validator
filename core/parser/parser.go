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

	var currentMethods Methods
	var currentMethod *Method
	var currentParameters []string

	for _, token := range scanner.tokens {
		switch token.type_ {
		case TOKEN_GROUP:
			currentMethods = Methods{}
			methodGroups = append(methodGroups, currentMethods)
		case TOKEN_METHOD:
			currentMethod = &Method{
				Name: token.value,
			}
			currentMethods = append(currentMethods, currentMethod)
			currentParameters = []string{}
		case TOKEN_PARAM_FLOAT, TOKEN_PARAM_INTEGER, TOKEN_PARAM_STRING:
			currentParameters = append(currentParameters, token.value)
		case TOKEN_ERROR:
			return nil, errors.New(token.value)
		}
	}

	return methodGroups, nil
}
