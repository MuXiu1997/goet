package templaterenderer

import (
	"bytes"
	"github.com/Masterminds/sprig/v3"
	"github.com/MuXiu1997/goet/pkg/templatefunc/toml"
	"github.com/MuXiu1997/goet/pkg/templatefunc/yaml"
	ttemplate "text/template"
)

// Render renders template with data
func Render(name string, templateContent string, data any) (string, error) {
	t, err := ttemplate.New(name).
		Funcs(sprig.TxtFuncMap()).
		Funcs(yaml.FuncMap()).
		Funcs(toml.FuncMap()).
		Parse(templateContent)
	if err != nil {
		return "", err
	}
	buffer := bytes.Buffer{}
	err = t.Execute(&buffer, data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
