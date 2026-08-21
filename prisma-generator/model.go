package prisma

import (
	"fmt"
	"strings"
)

type Model struct {
	PrismaName
	Fields []Field `json:"fields"`
}

type Field struct {
	PrismaName
	Type               string   `json:"type"`
	Kind               string   `json:"kind"`
	PK                 bool     `json:"isId"`
	List               bool     `json:"isList"`
	Required           bool     `json:"isRequired"`
	Unique             bool     `json:"isUnique"`
	RelationFromFields []string `json:"relationFromFields"`
	RelationToFields   []string `json:"relationToFields"`
}

func (f Field) Tags() (tags string) {
	if f.Kind == "object" {
		tags += f.relTags()
	} else {
		tags += f.dbTags()
	}

	return fmt.Sprintf("`bun:\"%s\"`", tags)
}

func (f Field) dbTags() string {
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

	return tags
}

func (f Field) relTags() string {
	tags := "rel:"

	if len(f.RelationFromFields) > 0 {
		tags += fmt.Sprintf("belongs-to,join:%s=%s", strings.Join(f.RelationFromFields, ","), strings.Join(f.RelationToFields, ","))
	} else if f.List {
		tags += "has-many"
	} else {
		tags += "has-one"
	}

	return tags
}

func (f Field) GoType() string {
	switch f.Type {
	case "Int":
		return "int"
	case "String":
		return "string"
	case "Boolean":
		return "bool"
	case "DateTime":
		return "time.Time"
	case "Decimal":
		return "float64"
	default:
		return f.Type
	}
}
