package file

import (
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func ReadFile(filePath string) ([]byte, error) {
	if strings.TrimSpace(filePath) == "-" {
		return io.ReadAll(os.Stdin)
	}
	u, err := url.Parse(filePath)
	if err != nil {
		return nil, err
	}

	switch u.Scheme {
	case "http", "https":
		resp, err := http.DefaultClient.Get(filePath)
		if err != nil {
			return nil, err
		}
		defer func(Body io.ReadCloser) {
			_ = Body.Close()
		}(resp.Body)
		return io.ReadAll(resp.Body)
	default:
		return os.ReadFile(filePath)
	}
}

func WriteFile(filePath string, content []byte) error {
	if strings.TrimSpace(filePath) == "" {
		_, err := os.Stdout.Write(content)
		return err
	}
	return os.WriteFile(filePath, content, 0o644) //nolint:gosec
}
