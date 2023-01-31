/*
The MIT License (MIT)

Copyright (c) 2018-2022 Tom Payne

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Code sourced from: Tom Payne, https://github.com/twpayne/chezmoi/blob/4cbdbd1820930a33914805d980979cf74ce0a75b/pkg/chezmoi/format.go
// Copyright belongs to Tom Payne
// Modified by MuXiu1997

package format

import (
	"encoding/json"
	"sort"
	"strings"

	"github.com/pelletier/go-toml/v2"
	goerrors "github.com/pkg/errors"
	"github.com/samber/lo"
	"gopkg.in/yaml.v3"
)

// Formats.
//
//goland:noinspection GoNameStartsWithPackageName
var (
	FormatJSON Format = formatJSON{}
	FormatTOML Format = formatTOML{}
	FormatYAML Format = formatYAML{}
)

// A Format is a serialization format.
type Format interface {
	Marshal(value any) ([]byte, error)
	Name() string
	Unmarshal(data []byte, value any) error
}

// A formatJSON implements the JSON serialization format.
type formatJSON struct{}

// A formatTOML implements the TOML serialization format.
type formatTOML struct{}

// A formatYAML implements the YAML serialization format.
type formatYAML struct{}

//goland:noinspection GoNameStartsWithPackageName
var (
	// FormatsByName is a map of all FormatsByName by name.
	FormatsByName = map[string]Format{
		"json": FormatJSON,
		"toml": FormatTOML,
		"yaml": FormatYAML,
	}

	FormatNames = func() []string {
		names := lo.Keys(FormatsByName)
		sort.Strings(names)
		return names
	}()

	// FormatsByExtension is a map of all Formats by extension.
	FormatsByExtension = map[string]Format{
		"json": FormatJSON,
		"toml": FormatTOML,
		"yaml": FormatYAML,
		"yml":  FormatYAML,
	}
)

// Marshal implements Format.Marshal.
func (formatJSON) Marshal(value any) ([]byte, error) {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return nil, goerrors.Wrap(err, "json")
	}
	return append(data, '\n'), nil
}

// Name implements Format.Name.
func (formatJSON) Name() string {
	return "json"
}

// Unmarshal implements Format.Unmarshal.
func (formatJSON) Unmarshal(data []byte, value any) error {
	return goerrors.Wrap(json.Unmarshal(data, value), "json")
}

// Marshal implements Format.Marshal.
func (formatTOML) Marshal(value any) ([]byte, error) {
	return toml.Marshal(value)
}

// Name implements Format.Name.
func (formatYAML) Name() string {
	return "yaml"
}

// Unmarshal implements Format.Unmarshal.
func (formatTOML) Unmarshal(data []byte, value any) error {
	return toml.Unmarshal(data, value)
}

// Marshal implements Format.Marshal.
func (formatYAML) Marshal(value any) ([]byte, error) {
	return yaml.Marshal(value)
}

// Name implements Format.Name.
func (formatTOML) Name() string {
	return "toml"
}

// Unmarshal implements Format.Unmarshal.
func (formatYAML) Unmarshal(data []byte, value any) error {
	return yaml.Unmarshal(data, value)
}

// TryUnmarshal tries to unmarshal data into value using all available formats.
// If the extension ext is non-empty, it is tried first.
func TryUnmarshal(data []byte, value any, ext string) error {
	var formatNameByExt string
	var errors []error
	formatByExt := FormatsByExtension[ext]
	if formatByExt != nil {
		formatNameByExt = formatByExt.Name()
		err := formatByExt.Unmarshal(data, value)
		if err == nil {
			return nil
		}
		errors = append(errors, err)
	}
	for _, formatName := range FormatNames {
		if formatName == formatNameByExt {
			continue
		}
		format := FormatsByName[formatName]
		err := format.Unmarshal(data, value)
		if err == nil {
			return nil
		}
		errors = append(errors, err)
	}
	return &TryUnmarshalError{errors}
}

type TryUnmarshalError struct {
	Errors []error
}

func (e *TryUnmarshalError) Error() string {
	var b strings.Builder
	b.WriteString("failed to unmarshal: \n")
	for _, err := range e.Errors {
		errLines := strings.Split(err.Error(), "\n")
		for i, errLine := range errLines {
			if i == 0 {
				b.WriteString(" -")
				b.WriteString(errLine)
			} else {
				b.WriteString("  ")
				b.WriteString(errLine)
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}
