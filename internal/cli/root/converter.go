package root

import (
	"encoding/json"
	"fmt"

	"github.com/abufattah/goku-cli/internal/platform/fileutil"
	"gopkg.in/yaml.v3"
)

func Convert(src []byte, inFormat, outFormat fileutil.Format) ([]byte, error) {
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

func decode(input []byte, format fileutil.Format, out any) error {
	switch format {
	case fileutil.FormatJSON:
		return json.Unmarshal(input, out)
	case fileutil.FormatYAML:
		return yaml.Unmarshal(input, out)
	default:
		return &fileutil.ErrUnsupportedFormat{Ext: string(format)}
	}
}

func encode(v any, format fileutil.Format) ([]byte, error) {
	switch format {
	case fileutil.FormatJSON:
		return json.MarshalIndent(v, "", "  ")
	case fileutil.FormatYAML:
		return yaml.Marshal(v)
	default:
		return nil, &fileutil.ErrUnsupportedFormat{Ext: string(format)}
	}
}
