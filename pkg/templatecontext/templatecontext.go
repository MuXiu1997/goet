package templatecontext

import (
	"github.com/MuXiu1997/goet/pkg/file"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

const (
	TemplateTypeFile = "file"
	TemplateTypeURL  = "url"
)

type TemplateContext struct {
	Template        string
	TemplateType    string
	TemplateContent string
	Output          string
	Values          any
}

func NewTemplateContext(template string, output string, values any) (*TemplateContext, error) {
	var (
		templatePath    string
		templateType    string
		templateContent []byte
		outputPath      string
		err             error
	)

	u, err := url.Parse(template)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		templatePath = template
		templateType = TemplateTypeURL
		templateContent, err = file.ReadURL(templatePath)
		if err != nil {
			return nil, err
		}
	default:
		if strings.TrimSpace(template) == "-" {
			templatePath = os.Stdin.Name()
		} else {
			templatePath, err = filepath.Abs(template)
			if err != nil {
				return nil, err
			}
		}
		templateType = TemplateTypeFile
		templateContent, err = file.ReadFile(templatePath)
		if err != nil {
			return nil, err
		}
	}

	if strings.TrimSpace(output) == "" {
		outputPath = os.Stdout.Name()
	} else {
		outputPath, err = filepath.Abs(output)
		if err != nil {
			return nil, err
		}
	}

	return &TemplateContext{
		Template:        templatePath,
		TemplateType:    templateType,
		TemplateContent: string(templateContent),
		Output:          outputPath,
		Values:          values,
	}, nil
}
