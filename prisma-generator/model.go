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
	DBName string  `json:"dbName"`
}

func (m Model) Table() string {
	if m.DBName != "" {
		return m.DBName
	}
	return m.Name
}

type Field struct {
	Name     FieldName `json:"name"`
	Type     FieldType `json:"type"`
	PK       bool      `json:"isId"`
	Required bool      `json:"isRequired"`
	Unique   bool      `json:"isUnique"`
	DBName   string    `json:"dbName"`
}

func (f Field) dbName() string {
	if f.DBName != "" {
		return f.DBName
	}
	return string(f.Name)
}

func (f Field) Tags() string {
	tags := f.dbName() + ",nullzero"

	if f.PK {
		tags += ",pk"
	}
	if f.Required {
		tags += ",notnull"
	}
	if f.Unique {
		tags += ",unique"
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
	case "DateTime":
		return "time.Time"
	default:
		return string(ft)
	}
}
