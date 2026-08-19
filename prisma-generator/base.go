package prisma

import "strings"

type PrismaName struct {
	Name     string `json:"name"`
	NameInDB string `json:"dbName"`
}

func (n PrismaName) DBName() string {
	if n.NameInDB != "" {
		return n.NameInDB
	}
	return n.Name
}

func (n PrismaName) GoName() string {
	parts := strings.Split(strings.ToLower(n.DBName()), "_")
	for i := 0; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			if val, exists := nameOverrides[parts[i]]; exists {
				parts[i] = val
			} else {
				parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
			}
		}
	}
	return strings.Join(parts, "")
}

var nameOverrides = map[string]string{
	"id": "ID",
}
