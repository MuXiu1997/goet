package toml

import (
	"github.com/MuXiu1997/goet/pkg/format"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"fromToml":     fromToml,
		"mustFromToml": mustFromToml,
		"toToml":       toToml,
		"mustToToml":   mustToToml,
	}
}

// fromToml decodes TOML into a structured value, ignoring errors.
func fromToml(data string) any {
	value, _ := mustFromToml(data)
	return value
}

// mustFromToml decodes TOML into a structured value, returning errors.
func mustFromToml(data string) (any, error) {
	var value any
	err := format.FormatTOML.Unmarshal([]byte(data), &value)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// toToml encodes an item into a TOML string, ignoring errors.
func toToml(value any) string {
	data, _ := mustToToml(value)
	return data
}

// mustToToml encodes an item into a TOML string, returning errors.
func mustToToml(value any) (string, error) {
	data, err := format.FormatTOML.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
