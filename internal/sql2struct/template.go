package sql2struct

import (
	"fmt"
	"os"
	"text/template"
	"tools/internal/word"
)

type StructTemplate struct {
	structTpl string
}

type StructTemplateColumn struct {
	Name    string
	Type    string
	Tag     string
	Comment string
}

type StructTemplateData struct {
	TableName string
	Columns   []*StructTemplateColumn
}

var structTpl string = `
type {{.TableName | ToCamelCase}} struct {
	{{range .Columns}}
	{{$len := len .Comment}}{{if gt $len 0}}// {{.Comment}}{{else}}// {{.Name}}{{end}}
	{{$typeLen := len .Type}}{{if gt $typeLen 0}}{{.Name | ToCamelCase}} {{.Type}} {{.Tag}} {{else}}{{.Name}}{{end}}
	{{end}}
}

func (model {{.TableName | ToCamelCase}}) TableName() string {
	return "{{.TableName}}"
}
`

func NewStructTemplate() *StructTemplate {
	return &StructTemplate{
		structTpl: structTpl,
	}
}

func (t *StructTemplate) MapColumns(tbCols []*TableColumn) []*StructTemplateColumn {
	tplCols := make([]*StructTemplateColumn, 0, len(tbCols))
	for _, col := range tbCols {
		tag := fmt.Sprintf("`"+"json:"+"\"%s\""+"`", col.ColumnName)
		tplCols = append(tplCols, &StructTemplateColumn{
			Name:    col.ColumnName,
			Tag:     tag,
			Type:    DBTypeToStructType[col.DataType],
			Comment: col.ColumnComment,
		})
	}
	return tplCols
}

func (t *StructTemplate) Generate(tableName string, tplCols []*StructTemplateColumn) error {
	tpl := template.Must(template.New("sql2struct").Funcs(
		template.FuncMap{
			"ToCamelCase": word.UnderscoreToUpperCamelCase,
		}).Parse(t.structTpl))

	tplData := StructTemplateData{
		TableName: tableName,
		Columns:   tplCols,
	}
	err := tpl.Execute(os.Stdout, tplData)
	if err != nil {
		return err
	}
	return nil
}
