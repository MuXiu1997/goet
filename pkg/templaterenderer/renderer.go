package templaterenderer

import (
	"bytes"
	"github.com/Masterminds/sprig/v3"
	"github.com/MuXiu1997/goet/pkg/templatefunc/osfspath"
	"github.com/MuXiu1997/goet/pkg/templatefunc/template"
	"github.com/MuXiu1997/goet/pkg/templatefunc/toml"
	"github.com/MuXiu1997/goet/pkg/templatefunc/yaml"
	ttemplate "text/template"
)
import tc "github.com/MuXiu1997/goet/pkg/templatecontext"

// Render renders template with data
func Render(templateContext *tc.TemplateContext) (string, error) {
	t, err := ttemplate.New(templateContext.Template).
		Funcs(sprig.TxtFuncMap()).
		Funcs(osfspath.FuncMap()).
		Funcs(yaml.FuncMap()).
		Funcs(toml.FuncMap()).
		Funcs(template.FuncMap(templateContext)).
		Parse(templateContext.TemplateContent)
	if err != nil {
		return "", err
	}
	buffer := bytes.Buffer{}
	err = t.Execute(&buffer, templateContext.Values)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
