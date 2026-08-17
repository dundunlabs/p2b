package prisma

type Output struct {
	Value string `json:"value"`
}

type Generator struct {
	Output Output `json:"output"`
}

type Model struct {
	Name string `json:"name"`
}

type Datamodel struct {
	Models []Model `json:"models"`
}

type DMMF struct {
	Datamodel Datamodel `json:"datamodel"`
}

type GenerateParams struct {
	Generator Generator `json:"generator"`
	DMMF      DMMF      `json:"dmmf"`
}
