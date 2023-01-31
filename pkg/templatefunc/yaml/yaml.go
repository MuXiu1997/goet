package yaml

import (
	"github.com/MuXiu1997/goet/pkg/format"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"fromYaml":     fromYaml,
		"mustFromYaml": mustFromYaml,
		"toYaml":       toYaml,
		"mustToYaml":   mustToYaml,
	}
}

// fromYaml decodes YAML into a structured value, ignoring errors.
func fromYaml(data string) any {
	value, _ := mustFromYaml(data)
	return value
}

// mustFromYaml decodes YAML into a structured value, returning errors.
func mustFromYaml(data string) (any, error) {
	var value any
	err := format.FormatYAML.Unmarshal([]byte(data), &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// toYaml encodes an item into a YAML string, ignoring errors.
func toYaml(value any) string {
	data, _ := mustToYaml(value)
	return data
}

// mustToYaml encodes an item into a YAML string, returning errors.
func mustToYaml(value any) (string, error) {
	data, err := format.FormatYAML.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
