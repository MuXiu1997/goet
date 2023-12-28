package template

import "path/filepath"
import ttemplate "text/template"

import tc "github.com/MuXiu1997/goet/pkg/templatecontext"

func FuncMap(templateContext *tc.TemplateContext) ttemplate.FuncMap {
	return ttemplate.FuncMap{
		"templateFilePath": templateFilePath(templateContext),
		"templateDirPath":  templateDirPath(templateContext),
		"outputFilePath":   outputFilePath(templateContext),
		"outputDirPath":    outputDirPath(templateContext),
	}
}

// templateFilePath returns the path of the template file.
// If the template is stdin, "/dev/stdin" is returned.
// If the template is a URL, the URL is returned.
func templateFilePath(templateContext *tc.TemplateContext) func() string {
	return func() string {
		return templateContext.Template
	}
}

// templateDirPath returns the directory of the template file.
// If the template is stdin, "/dev" is returned.
// If the template is a URL, "" is returned.
func templateDirPath(templateContext *tc.TemplateContext) func() string {
	return func() string {
		if templateContext.TemplateType == tc.TemplateTypeFile {
			return filepath.Dir(templateContext.Template)
		} else {
			return ""
		}
	}
}

// outputFilePath returns the path of the output file.
// If the output is stdout, "/dev/stdout" is returned.
func outputFilePath(templateContext *tc.TemplateContext) func() string {
	return func() string {
		return templateContext.Output
	}
}

// outputDirPath returns the directory of the output file.
// If the output is stdout, "/dev" is returned.
func outputDirPath(templateContext *tc.TemplateContext) func() string {
	return func() string {
		return filepath.Dir(templateContext.Output)
	}
}
