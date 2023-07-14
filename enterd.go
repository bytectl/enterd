package enterd

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"
	"math"
	"os"
	"path/filepath"
	"strings"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
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
	// g.Nodes[0].Fields[0].Type.Type.ConstName()
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

func defaultSize(c *schema.Column) int64 {
	size := int64(255)
	return size
}

func ERDType(c *schema.Column) (t string) {
	if c.SchemaType != nil && c.SchemaType[dialect.MySQL] != "" {
		// MySQL returns the column type lower cased.
		return strings.ToLower(c.SchemaType[dialect.MySQL])
	}
	switch c.Type {
	case field.TypeBool:
		t = "boolean"
	case field.TypeInt8:
		t = "tinyint"
	case field.TypeUint8:
		t = "tinyint unsigned"
	case field.TypeInt16:
		t = "smallint"
	case field.TypeUint16:
		t = "smallint unsigned"
	case field.TypeInt32:
		t = "int"
	case field.TypeUint32:
		t = "int unsigned"
	case field.TypeInt, field.TypeInt64:
		t = "bigint"
	case field.TypeUint, field.TypeUint64:
		t = "bigint unsigned"
	case field.TypeBytes:
		size := int64(math.MaxUint16)
		if c.Size > 0 {
			size = c.Size
		}
		switch {
		case size <= math.MaxUint8:
			t = "tinyblob"
		case size <= math.MaxUint16:
			t = "blob"
		case size < 1<<24:
			t = "mediumblob"
		case size <= math.MaxUint32:
			t = "longblob"
		}
	case field.TypeJSON:
		t = "json"
	case field.TypeString:
		size := c.Size
		if size == 0 {
			size = defaultSize(c)
		}
		switch {
		case size <= math.MaxUint16:
			t = fmt.Sprintf("varchar(%d)", size)
		case size == 1<<24-1:
			t = "mediumtext"
		default:
			t = "longtext"
		}
	case field.TypeFloat32, field.TypeFloat64:
		t = "double"
	case field.TypeTime:
		t = "timestamp"
	case field.TypeEnum:
		values := make([]string, len(c.Enums))
		for i, e := range c.Enums {
			values[i] = fmt.Sprintf("'%s'", e)
		}
		t = fmt.Sprintf("enum(%s)", strings.Join(values, ", "))
	case field.TypeUUID:
		t = "char(36) binary"
	case field.TypeOther:
		t = c.Name
	default:
		panic(fmt.Sprintf("unsupported type %q for column %q", c.Type.String(), c.Name))
	}
	return t
}
