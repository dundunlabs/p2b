package prisma

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Model struct {
	Name   string  `json:"name"`
	Fields []Field `json:"fields"`
}

type Field struct {
	Name     FieldName `json:"name"`
	Type     FieldType `json:"type"`
	PK       bool      `json:"isId"`
	Required bool      `json:"isRequired"`
	Unique   bool      `json:"isUnique"`
}

func (f Field) Tags() (tags string) {
	if f.PK {
		tags += ",pk"
	}
	if f.Required {
		tags += ",notnull"
	}
	if f.Unique {
		tags += ",unique"
	}

	if tags == "" {
		return ""
	}
	return fmt.Sprintf("`bun:\"%s\"`", tags)
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
