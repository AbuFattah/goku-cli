package root

import (
	"fmt"
	"path/filepath"
	"strings"
)

type Format string

const (
	FormatJSON Format = "json"
	FormatYAML Format = "yaml"
)

type ErrUnsupportedFormat struct {
	Ext string
}

func (e *ErrUnsupportedFormat) Error() string {
	return fmt.Sprintf("unsupported format: %s", e.Ext)
}

func DetectFormat(path string) (Format, error) {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".json":
		return FormatJSON, nil
	case ".yaml", ".yml":
		return FormatYAML, nil
	default:
		return "", &ErrUnsupportedFormat{Ext: ext}
	}
}
