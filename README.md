# Goet 🖨️

single-executable template renderer, powered by go template, sprig.

## Usage

```help
single-executable template renderer, powered by go template, sprig.

Usage:
  goet [flags]

Flags:
  -t, --template string          specify template file, "-" or unset means stdin (default "-")
  -o, --output string            specify output file, unset means stdout
  -f, --values strings           specify values in initConfig JSON/TOML/YAML file or initConfig URL (can specify multiple)
  -J, --set-json stringArray     set JSON values on the command line (can specify multiple or separate values with commas: key1=jsonval1,key2=jsonval2)
  -s, --set stringArray          set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)
  -S, --set-string stringArray   set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)
  -F, --set-file stringArray     set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)
  -h, --help                     help for goet
  -v, --version                  version for goet
```

## Example

```bash
$ goet -s name=MuXiu1997 <<EOF
Hello {{ .name }}, you're currently in {{ env "PWD" }}
EOF

# output:
Hello MuXiu1997, you're currently in /Users/muxiu1997/Projects/goet
```

## Install

prebuilt binaries are available on [release page](https://github.com/MuXiu1997/goet/releases)

## Template Functions

### Sprig

- [Sprig Function Documentation](http://masterminds.github.io/sprig/)

### OS/FS/PATH

- `lookPath` - lookPath runs exec.LookPath on name and returns the path or an empty string if the name is not found.
- `lstat` - lstat runs os.Lstat on name and returns structured data that contains the fields `name`, `size`, `mode`, `perm`, `modTime`, `isDir`, and `type` if name exists.
- `readFile` - readFile runs os.ReadFile on name and returns the contents.
- `cmdOutput` - cmdOutput runs exec.Command on name and args and returns the output.
- `glob` - glob runs filepath.Glob on pattern and returns the matches.
- `joinPath` - joinPath runs filepath.Join on elements and returns the joined path with the OS-specific path separator.

### YAML

- `fromYaml` - fromYaml decodes YAML into a structured value, ignoring errors.
- `mustFromYaml` - mustFromYaml decodes YAML into a structured value, returning errors.
- `toYaml` - toYaml encodes an item into a YAML string, ignoring errors.
- `mustToYaml` - mustToYaml encodes an item into a YAML string, returning errors.

### TOML

- `fromToml` - fromToml decodes TOML into a structured value, ignoring errors.
- `mustFromToml` - mustFromToml decodes TOML into a structured value, returning errors.
- `toToml` - toToml encodes an item into a TOML string, ignoring errors.
- `mustToToml` - mustToToml encodes an item into a TOML string, returning errors.

## Related
- sprig template functions - [Masterminds/sprig](https://github.com/Masterminds/sprig)
- additional template functions reference - [twpayne/chezmoi](https://github.com/twpayne/chezmoi)

## License

[MIT](./LICENSE)