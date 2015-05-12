package parser

import (
	"fmt"
)

const eof = -1

type tokenType int

type token struct {
	type_    tokenType
	position int
	value    string
}

const (
	TOKEN_NONE tokenType = iota

	TOKEN_ERROR
	TOKEN_EOF

	TOKEN_GROUP
	TOKEN_METHOD

	TOKEN_PARAM_INTEGER
	TOKEN_PARAM_FLOAT
	TOKEN_PARAM_STRING
)

func (this token) String() string {
	return fmt.Sprintf("<type=%d, value='%s', pos=%d>", this.type_, this.value, this.position)
}
