package prisma

import (
	"encoding/json"
	"fmt"
	"strings"
)

var models = map[string]map[string]*Field{}

type Model struct {
	PrismaName
	Fields     []*Field    `json:"fields"`
	PrimaryKey *PrimaryKey `json:"primaryKey"`
}

func (m *Model) UnmarshalJSON(data []byte) error {
	type Alias Model
	a := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	if err := json.Unmarshal(data, a); err != nil {
		return err
	}

	model := map[string]*Field{}
	for _, f := range m.Fields {
		f.m = m
		model[f.Name] = f
	}
	models[m.Name] = model

	if m.PrimaryKey != nil {
		for _, f := range m.PrimaryKey.Fields {
			if mf, ok := model[f]; ok {
				mf.PK = true
			}
		}
	}

	return nil
}

type Field struct {
	m *Model

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
	tags := f.DBName()

	if f.PK {
		tags += ",pk"
	} else if f.Required {
		tags += ",notnull"
	}
	if f.Unique {
		tags += ",unique"
	}

	return tags + ",nullzero"
}

func (f Field) relTags() string {
	tags := "rel:"

	if len(f.RelationFromFields) > 0 {
		var ff, tf []string
		for _, n := range f.RelationFromFields {
			ff = append(ff, models[f.m.Name][n].DBName())
		}
		for _, n := range f.RelationToFields {
			tf = append(tf, models[f.Type][n].DBName())
		}
		tags += fmt.Sprintf("belongs-to,join:%s=%s", strings.Join(ff, ","), strings.Join(tf, ","))
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

type PrimaryKey struct {
	Fields []string `json:"fields"`
}
