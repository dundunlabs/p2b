package generator

import (
	"bytes"
	_ "embed"
	"os"
	"path/filepath"
	"text/template"

	"github.com/dundunlabs/p2b/prisma-generator"
	"golang.org/x/tools/imports"
)

//go:embed templates/models.gotpl
var modelsGoTpl string

const filename = "models_gen.go"

var (
	modelsTpl = template.Must(template.New("models").Parse(modelsGoTpl))
)

type Generator struct {
	params prisma.GenerateParams
}

func New(params prisma.GenerateParams) *Generator {
	return &Generator{params}
}

func (g *Generator) Generate() error {
	dir := g.params.Generator.Output.Value
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	output := filepath.Join(dir, filename)
	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	var data bytes.Buffer
	if err := modelsTpl.Execute(&data, g.params.DMMF); err != nil {
		return err
	}

	content, err := imports.Process(output, data.Bytes(), nil)
	if err != nil {
		return err
	}

	_, err = file.Write(content)
	return err
}
