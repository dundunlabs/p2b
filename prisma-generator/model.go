package prisma

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

type Model struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name FieldName `json:"name"`
	Type FieldType `json:"type"`
}

type FieldName string

func (fn FieldName) String() string {
	s := string(fn)

	// 1. Check for exact abbreviation matches
	if val, exists := nameOverrides[strings.ToLower(s)]; exists {
		return val
	}

	// 2. Fallback to standard first-letter capitalization
	r, size := utf8.DecodeRuneInString(s)
	if r == utf8.RuneError {
		return s
	}
	return string(unicode.ToUpper(r)) + s[size:]
}

var nameOverrides = map[string]string{
	"id": "ID",
}

type FieldType string

func (ft FieldType) String() string {
	switch ft {
	case "Int":
		return "int"
	case "String":
		return "string"
	default:
		return string(ft)
	}
}
