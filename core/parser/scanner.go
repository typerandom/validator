package parser

import (
	"fmt"
	"unicode/utf8"
)

type scanner struct {
	value string

	start    int
	position int
	width    int

	tokens []*token
}

func (this *scanner) next() rune {
	if int(this.position) >= len(this.value) {
		this.width = 0
		return eof
	}

	char, width := utf8.DecodeRuneInString(this.value[this.position:])
	this.width = width
	this.position += this.width

	return char
}

func (this *scanner) peek() rune {
	char := this.next()
	this.backup()
	return char
}

func (this *scanner) backup() {
	this.position -= this.width
}

func (this *scanner) skip() {
	this.start = this.position
}

func (this *scanner) text() string {
	return this.value[this.start:this.position]
}

func (this *scanner) length() int {
	return this.position - this.start
}

func (this *scanner) emit(token tokenType) {
	this.emitValue(token, this.text())
}

func (this *scanner) emitValue(typ tokenType, value string) {
	this.tokens = append(this.tokens, &token{
		type_:    typ,
		position: this.start,
		value:    value,
	})
	this.skip()
}

func (this *scanner) errorf(format string, args ...interface{}) lexer {
	this.tokens = append(this.tokens, &token{
		type_:    TOKEN_ERROR,
		position: this.start,
		value:    fmt.Sprintf(format, args...),
	})
	return nil
}

func (this *scanner) unexpectedCharError() lexer {
	return this.errorf("Unexpected character %#U at position %d.", this.value[this.position-1], this.position)
}

func (this *scanner) UnexpectedEndError() lexer {
	return this.errorf("Unexpected end at position %d.", this.position)
}
