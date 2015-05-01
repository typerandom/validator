package core

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

const (
	STATE_NAME              = 0
	STATE_OPTION            = 1
	STATE_OPTION_ESCAPE     = 2
	STATE_OPTION_WHITESPACE = 3
)

type Tag struct {
	Name    string
	Options []string
}

func (this *Tag) String() string {
	serializedOptions := strings.Join(this.Options, "', '")

	if len(serializedOptions) == 0 {
		serializedOptions = "(none)"
	} else {
		serializedOptions = "'" + serializedOptions + "'"
	}

	return "{ name: '" + this.Name + "', options: " + serializedOptions + " }"
}

type TagGroup []*Tag

func (this TagGroup) String() string {
	result := ""

	for _, tag := range this {
		if result != "" {
			result += ", "
		}
		result += tag.String()
	}

	return result
}

func parseTag(rawTag string) ([]TagGroup, error) {
	var state int
	var buffer bytes.Buffer
	var tagBuffer TagGroup

	var groups []TagGroup
	var tag *Tag = &Tag{}

	getParserError := func(offset int, char rune) error {
		return errors.New(fmt.Sprintf("Parser error: Unexpected character '%c' at position %d.", char, offset))
	}

	//fmt.Println("Offset	Char	State")

	//stateNames := []string{"name", "option", "escape", "whitespace"}

	for offset, char := range rawTag {
		//fmt.Printf("%d	%c	%s\n", offset, char, stateNames[state])
		switch state {
		case STATE_NAME:
			switch char {
			case '(':
				if buffer.Len() == 0 {
					return nil, getParserError(offset, char)
				}
				state = STATE_OPTION
				tag.Name = buffer.String()
				buffer.Reset()
			case '|':
				if buffer.Len() != 0 {
					tag.Name = buffer.String()
					tagBuffer = append(tagBuffer, tag)
					buffer.Reset()
					tag = &Tag{}
				}
				if len(tagBuffer) != 0 {
					groups = append(groups, tagBuffer)
					tagBuffer = make([]*Tag, 0)
				}
			case ',':
				if buffer.Len() > 0 {
					tag.Name = buffer.String()
					buffer.Reset()
				}

				if len(tag.Name) > 0 {
					tagBuffer = append(tagBuffer, tag)
				}

				tag = &Tag{}
			default:
				goto WRITE_CHAR
			}
			continue

		case STATE_OPTION:
			switch char {
			case '\\':
				state = STATE_OPTION_ESCAPE
			case ')':
				state = STATE_NAME
				if buffer.Len() > 0 {
					tag.Options = append(tag.Options, buffer.String())
					tagBuffer = append(tagBuffer, tag)
					tag = &Tag{}
					buffer.Reset()
				}
			case ',':
				if buffer.Len() > 0 {
					state = STATE_OPTION_WHITESPACE
					tag.Options = append(tag.Options, buffer.String())
					buffer.Reset()
				}
			default:
				goto WRITE_CHAR
			}
			continue

		case STATE_OPTION_WHITESPACE:
			state = STATE_OPTION

			if char == ' ' || char == '	' {
				continue
			}

		case STATE_OPTION_ESCAPE:
			state = STATE_OPTION
		}

	WRITE_CHAR:
		buffer.WriteRune(char)
	}

	if buffer.Len() > 0 {
		tag.Name = buffer.String()
	}

	if tag != nil && len(tag.Name) > 0 {
		tagBuffer = append(tagBuffer, tag)
	}

	if len(tagBuffer) > 0 {
		groups = append(groups, tagBuffer)
	}

	/*fmt.Println(rawTag)

	for _, group := range groups {
		fmt.Println(group)
	}*/

	return groups, nil
}
