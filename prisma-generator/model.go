package prisma

import (
	"fmt"
)

type Model struct {
	PrismaName
	Fields []Field `json:"fields"`
}

type Field struct {
	PrismaName
	Type     FieldType `json:"type"`
	PK       bool      `json:"isId"`
	Required bool      `json:"isRequired"`
	Unique   bool      `json:"isUnique"`
}

func (f Field) Tags() string {
	tags := f.DBName() + ",nullzero"

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
