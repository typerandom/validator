package parser

import (
	"bytes"
)

type lexer func(*scanner) lexer

func isAlpha(char rune) bool {
	return char >= 'a' && char <= 'z'
}

func isNumeric(char rune) bool {
	return char >= '0' && char <= '9'
}

func isAlphaNumeric(char rune) bool {
	return isAlpha(char) || isNumeric(char)
}

func isWhiteSpace(char rune) bool {
	return char == ' ' || char == 9 // 9 = tab
}

func lexWhiteSpace(scanner *scanner, returnTo lexer) lexer {
	for {
		char := scanner.next()
		if !isWhiteSpace(char) || char == eof {
			break
		}
	}

	scanner.backup()
	scanner.skip()

	return returnTo
}

func skipIndexesOfString(value string, indexesToSkip []int) string {
	if len(value) == 0 || len(indexesToSkip) == 0 {
		return value
	}

	var offset int
	var buffer bytes.Buffer
	var totalSkipIndexes int = len(indexesToSkip)

	for i, char := range value {
		if offset < totalSkipIndexes && indexesToSkip[offset] == i {
			offset++
		} else {
			buffer.WriteRune(char)
		}
	}

	return buffer.String()
}

func lexParamValueText(scanner *scanner) lexer {
	var escapes []int

TEXT_SCAN:
	for {
		switch scanner.next() {
		case '\\':
			escapes = append(escapes, scanner.position-scanner.start-1)
			scanner.next()
		case '´':
			scanner.backup()
			break TEXT_SCAN
		case eof:
			return scanner.UnexpectedEndError()
		}
	}

	textValue := scanner.text()

	if len(escapes) > 0 {
		textValue = skipIndexesOfString(textValue, escapes)
	}

	scanner.emitValue(TOKEN_PARAM_STRING, textValue)

	scanner.next()
	scanner.skip()

	return lexParams
}

func lexParamValueNumber(scanner *scanner) lexer {
	var returnTo lexer
	isFloat := false

NUMBER_SCAN:
	for {
		switch char := scanner.next(); {
		case char == '+' || char == '-':
			if scanner.length() != 1 {
				return scanner.unexpectedCharError()
			}
		case isNumeric(char):
			continue
		case char == '.':
			if scanner.length() == 1 || isFloat {
				return scanner.unexpectedCharError()
			}
			isFloat = true
		case char == ',':
			returnTo = lexParams
			break NUMBER_SCAN
		case char == ')':
			returnTo = lexParams
			break NUMBER_SCAN
		case char == eof:
			return scanner.UnexpectedEndError()
		default:
			return scanner.unexpectedCharError()
		}
	}

	scanner.backup()

	if isFloat {
		scanner.emit(TOKEN_PARAM_FLOAT)
	} else {
		scanner.emit(TOKEN_PARAM_INTEGER)
	}

	return returnTo
}

func lexParamValue(scanner *scanner) lexer {
	switch char := scanner.next(); {
	case char == '+' || char == '-' || isNumeric(char):
		scanner.backup()
		return lexParamValueNumber
	case isWhiteSpace(char):
		return lexWhiteSpace(scanner, lexParamValue)
	case char == '´':
		scanner.skip()
		return lexParamValueText
	default:
		return scanner.unexpectedCharError()
	}
}

func lexParams(scanner *scanner) lexer {
	switch char := scanner.next(); {
	case char == ',':
		scanner.skip()
		return lexParamValue
	case char == '(':
		var returnTo lexer = lexParamValue

		if scanner.peek() == ')' {
			scanner.next()
			scanner.skip()
			returnTo = lexMethod
		}

		scanner.skip()
		return returnTo
	case char == ')':
		scanner.skip()
		return lexMethod
	default:
		return scanner.unexpectedCharError()
	}
}

func lexMethodName(scanner *scanner) lexer {
	var returnTo lexer

NAME_SCAN:
	for {
		switch char := scanner.next(); {
		case isAlphaNumeric(char) || char == '_':
			continue
		case char == '|':
			returnTo = lexGroup
			break NAME_SCAN
		case char == ',':
			returnTo = lexMethod
			break NAME_SCAN
		case char == '(':
			returnTo = lexParams
			break NAME_SCAN
		case char == eof:
			break NAME_SCAN
		default:
			return scanner.unexpectedCharError()
		}
	}

	scanner.backup()
	scanner.emit(TOKEN_METHOD)

	return returnTo
}

func lexMethod(scanner *scanner) lexer {
	switch char := scanner.next(); {
	case isAlphaNumeric(char) || char == '_':
		scanner.backup()
		return lexMethodName
	case char == '|':
		scanner.backup()
		return lexGroup
	case char == ',':
		scanner.skip()
		return lexMethod
	case char == '(':
		scanner.skip()
		return lexParams
	case char == eof:
		return nil
	default:
		return scanner.unexpectedCharError()
	}
}

func lexGroup(scanner *scanner) lexer {
	switch char := scanner.next(); {
	case isAlphaNumeric(char) || char == '_':
		scanner.backup()
		return lexMethod
	case char == '|':
		next := scanner.peek()
		if scanner.position == 1 || next == '|' || next == eof {
			return scanner.unexpectedCharError()
		}
		scanner.emit(TOKEN_GROUP)
		return lexMethod
	default:
		return scanner.unexpectedCharError()
	}
}
