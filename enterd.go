package enterd

import (
	"bytes"
	_ "embed"
	"html/template"
	"os"
	"path/filepath"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

var (
	//go:embed erd.go.tmpl
	tmplerd string
	erdtmpl = template.Must(
		template.New("erd").
			Funcs(template.FuncMap{
				"ERDType": ERDType,
			}).Parse(tmplerd),
	)
)

func generateERD(g *gen.Graph) ([]byte, error) {
	var b bytes.Buffer
	if err := erdtmpl.Execute(&b, g); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// VisualizeSchema is an ent hook that generates a static html page that visualizes the schema graph.
func VisualizeSchema(next gen.Generator) gen.Generator {
	return gen.GenerateFunc(func(g *gen.Graph) error {
		if err := next.Generate(g); err != nil {
			return err
		}
		buf, err := generateERD(g)
		if err != nil {
			return err
		}
		path := filepath.Join(g.Config.Target, "schema-erd.md")
		return os.WriteFile(path, buf, 0644)
	})
}

type Extension struct {
	entc.DefaultExtension
}

func (Extension) Hooks() []gen.Hook {
	return []gen.Hook{
		VisualizeSchema,
	}
}

func (Extension) Templates() []*gen.Template {
	return []*gen.Template{
		//gen.MustParse(gen.NewTemplate("enterd").Parse(tmplfile)),
	}
}

func GenerateMMD(schemaPath string, cfg *gen.Config) ([]byte, error) {
	g, err := entc.LoadGraph(schemaPath, cfg)
	if err != nil {
		return nil, err
	}
	return generateERD(g)
}

func ERDType(t string) string {
	switch t {
	case "[]byte":
		return "bytes"
	case "int8", "int16":
		return "int32"
	case "int":
		return "int64"
	case "float64":
		return "double"
	case "float32":
		return "float"
	case "time.Time":
		return "datetime"
	}
	return t
}
