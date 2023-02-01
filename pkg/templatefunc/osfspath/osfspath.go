package osfspath

import (
	"errors"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func FuncMap() template.FuncMap {
	return template.FuncMap{
		"lookPath":  lookPath,
		"lstat":     lstat,
		"readFile":  readFile,
		"cmdOutput": cmdOutput,
		"glob":      glob,
		"joinPath":  joinPath,
	}
}

// lookPath runs exec.LookPath on name and returns the path or an empty string if the name is not found.
func lookPath(name string) (string, error) {
	path, err := exec.LookPath(name)
	if errors.Is(err, exec.ErrNotFound) || errors.Is(err, fs.ErrNotExist) {
		return "", nil
	}
	return path, err
}

// lstat runs os.Lstat on name and returns structured data that contains the fields `name`, `size`, `mode`, `perm`, `modTime`, `isDir`, and `type` if name exists.
func lstat(name string) (map[string]any, error) {
	fileInfo, err := os.Lstat(name)
	if err == nil {
		return fileInfoToMap(fileInfo), nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return nil, nil
	}
	return nil, err
}

// readFile runs os.ReadFile on name and returns the contents.
func readFile(name string) (string, error) {
	contents, err := os.ReadFile(name)
	if err != nil {
		return "", err
	}
	return string(contents), nil
}

// cmdOutput runs exec.Command on name and args and returns the output.
func cmdOutput(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// glob runs filepath.Glob on pattern and returns the matches.
func glob(pattern string) []string {
	matches, _ := filepath.Glob(pattern)
	return matches
}

// joinPath runs filepath.Join on elements and returns the joined path with the OS-specific path separator.
func joinPath(elements ...string) string {
	return filepath.Join(elements...)
}
