package core

import (
	"bytes"
	"strings"
)

const (
	StateName             = 0
	StateOption           = 1
	StateOptionEscape     = 2
	StateOptionWhitespace = 3
)

type Tag struct {
	Name    string
	Options []string
}

func (this *Tag) String() string {
	serializedOptions := strings.Join(this.Options, ", ")

	if len(serializedOptions) == 0 {
		serializedOptions = "(none)"
	} else {
		serializedOptions = "'" + serializedOptions + "'"
	}

	return "{ name: '" + this.Name + "', options: " + serializedOptions + " }"
}

func parseTag(rawTag string) []*Tag {
	var tags []*Tag

	var buffer bytes.Buffer
	var currentTag *Tag
	var currentState int

	for _, char := range rawTag {
		switch {
		case currentState == StateOptionEscape:
			buffer.WriteRune(char)
			currentState = StateOption
		case currentState == StateOption && char == '\\':
			currentState = StateOptionEscape
		case char == '(':
			currentState = StateOption
			currentTag.Name = buffer.String()
			buffer.Reset()
		case currentState == StateOption && char == ')':
			currentState = StateName
			if buffer.Len() > 0 {
				currentTag.Options = append(currentTag.Options, buffer.String())
				buffer.Reset()
			}
		case currentState == StateOptionWhitespace:
			if char != ' ' && char != '	' {
				currentState = StateOption
				buffer.WriteRune(char)
			}
		case currentState == StateOption && char == ',':
			if buffer.Len() > 0 {
				currentState = StateOptionWhitespace
				currentTag.Options = append(currentTag.Options, buffer.String())
				buffer.Reset()
			}
		case currentState == StateName && char == ',':
			if buffer.Len() > 0 {
				currentTag.Name = buffer.String()
				buffer.Reset()
			}

			if len(currentTag.Name) > 0 {
				tags = append(tags, currentTag)
			}

			currentTag = &Tag{}
		default:
			if currentTag == nil {
				currentTag = &Tag{}
			}
			buffer.WriteRune(char)
		}
	}

	if buffer.Len() > 0 {
		currentTag.Name = buffer.String()
	}

	if currentTag != nil && len(currentTag.Name) > 0 {
		tags = append(tags, currentTag)
	}

	return tags
}
