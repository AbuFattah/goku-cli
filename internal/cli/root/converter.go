package root

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
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

func Convert(src []byte, inFormat, outFormat Format) ([]byte, error) {
	if inFormat == outFormat {
		return nil, fmt.Errorf("input filetype and output filetype should be different")
	}

	var temp any
	if err := decode(src, inFormat, &temp); err != nil {
		return nil, fmt.Errorf("decoding %s: %w", inFormat, err)
	}

	out, err := encode(temp, outFormat)
	if err != nil {
		return nil, fmt.Errorf("encoding %s: %w", outFormat, err)
	}

	return out, nil
}

func decode(input []byte, format Format, out any) error {
	switch format {
	case FormatJSON:
		return json.Unmarshal(input, out)
	case FormatYAML:
		return yaml.Unmarshal(input, out)
	default:
		return &ErrUnsupportedFormat{Ext: string(format)}
	}
}

func encode(v any, format Format) ([]byte, error) {
	switch format {
	case FormatJSON:
		return json.MarshalIndent(v, "", "  ")
	case FormatYAML:
		return yaml.Marshal(v)
	default:
		return nil, &ErrUnsupportedFormat{Ext: string(format)}
	}
}
