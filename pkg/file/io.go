package file

import (
	"io"
	"net/http"
	"os"
	"strings"
)

func ReadFile(filePath string) ([]byte, error) {
	if strings.TrimSpace(filePath) == "-" {
		return io.ReadAll(os.Stdin)
	}
	return os.ReadFile(filePath)
}

func ReadURL(url string) ([]byte, error) {
	resp, err := http.DefaultClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	return io.ReadAll(resp.Body)
}

func WriteFile(filePath string, content []byte) error {
	if strings.TrimSpace(filePath) == "" {
		_, err := os.Stdout.Write(content)
		return err
	}
	return os.WriteFile(filePath, content, 0o644) //nolint:gosec
}
