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

	TOKEN_ARG_INTEGER
	TOKEN_ARG_FLOAT
	TOKEN_ARG_STRING
	TOKEN_ARG_BOOLEAN
	TOKEN_ARG_NIL
)

func (this token) String() string {
	return fmt.Sprintf("<type=%d, value='%s', pos=%d>", this.type_, this.value, this.position)
}
