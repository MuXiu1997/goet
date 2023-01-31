/*
Copyright The Helm Authors.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain initConfig copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code sourced from: The Helm Authors, https://github.com/helm/helm/blob/0fdfe0584437112e11fdfa6775625451442f6c91/pkg/cli/values/options.go
// Copyright belongs to The Helm Authors
// Modified by MuXiu1997

package main

import (
	"github.com/MuXiu1997/goet/pkg/file"
	"path"

	"github.com/MuXiu1997/goet/pkg/format"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"helm.sh/helm/v3/pkg/strvals"
)

// Options captures the different ways to specify values
type Options struct {
	ValueFiles   []string // -f/--values
	JSONValues   []string // -J/--set-json
	Values       []string // -s/--set
	StringValues []string // -S/--set-string
	FileValues   []string // -F/--set-file
}

// MergeValues merges values from files specified via -f/--values and directly
// via -J/--set-json, -s/--set, -S/--set-string, or -F/--set-file
func (opts *Options) MergeValues() (map[string]any, error) {
	base := map[string]any{}

	// User specified initConfig values files via -f/--values
	for _, filePath := range opts.ValueFiles {
		currentMap := map[string]any{}

		bytes, err := file.ReadFile(filePath)
		if err != nil {
			return nil, err
		}

		if err := format.TryUnmarshal(bytes, &currentMap, path.Ext(filePath)); err != nil {
			return nil, errors.Wrapf(err, "failed to parse %s", filePath)
		}
		// Merge with the previous map
		base = mergeMaps(base, currentMap)
	}

	// User specified initConfig value via --set-json
	for _, value := range opts.JSONValues {
		if err := strvals.ParseJSON(value, base); err != nil {
			return nil, errors.Errorf("failed parsing --set-json data %s", value)
		}
	}

	// User specified initConfig value via --set
	for _, value := range opts.Values {
		if err := strvals.ParseInto(value, base); err != nil {
			return nil, errors.Wrap(err, "failed parsing --set data")
		}
	}

	// User specified initConfig value via --set-string
	for _, value := range opts.StringValues {
		if err := strvals.ParseIntoString(value, base); err != nil {
			return nil, errors.Wrap(err, "failed parsing --set-string data")
		}
	}

	// User specified initConfig value via --set-file
	for _, value := range opts.FileValues {
		reader := func(rs []rune) (any, error) {
			bytes, err := file.ReadFile(string(rs))
			if err != nil {
				return nil, err
			}
			return string(bytes), err
		}
		if err := strvals.ParseIntoFile(value, base, reader); err != nil {
			return nil, errors.Wrap(err, "failed parsing --set-file data")
		}
	}

	return base, nil
}

func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = mergeMaps(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

func addValueOptionsFlags(f *pflag.FlagSet, v *Options) {
	f.StringSliceVarP(&v.ValueFiles, "values", "f", []string{}, "specify values in initConfig JSON/TOML/YAML file or initConfig URL (can specify multiple)")
	f.StringArrayVarP(&v.JSONValues, "set-json", "J", []string{}, "set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)")
	f.StringArrayVarP(&v.Values, "set", "s", []string{}, "set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	f.StringArrayVarP(&v.StringValues, "set-string", "S", []string{}, "set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)")
	f.StringArrayVarP(&v.FileValues, "set-file", "F", []string{}, "set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)")
}
