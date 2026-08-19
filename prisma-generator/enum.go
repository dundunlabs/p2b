package prisma

type Enum struct {
	PrismaName
	Values []PrismaName `json:"values"`
}
