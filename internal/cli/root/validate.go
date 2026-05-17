package root

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func validateFilepathFlags(input, output string) error {

	if input == "" {
		return fmt.Errorf("input file is required (-i/--input)")
	}
	if output == "" {
		return fmt.Errorf("output file is required (-o/--output)")
	}

	if _, err := os.Stat(input); os.IsNotExist(err) {
		return fmt.Errorf("input file does not exist: %s", input)
	}

	info, _ := os.Stat(input)
	if info.IsDir() {
		return fmt.Errorf("input path is a directory, expected a file: %s", input)
	}

	allowed := map[string]bool{".json": true, ".yaml": true, ".yml": true}
	inFormat := strings.ToLower(filepath.Ext(input))
	outFormat := strings.ToLower(filepath.Ext(output))

	if !allowed[inFormat] {
		return fmt.Errorf("unsupported input format %q, allowed: json, yaml", inFormat)
	}
	if !allowed[outFormat] {
		return fmt.Errorf("unsupported output format %q, allowed: json, yaml", outFormat)
	}

	if inFormat == outFormat {
		return fmt.Errorf("input and output formats should be different")
	}

	outDir := filepath.Dir(output)
	if _, err := os.Stat(outDir); os.IsNotExist(err) {
		return fmt.Errorf("output directory does not exist: %s", outDir)
	}

	absIn, _ := filepath.Abs(input)
	absOut, _ := filepath.Abs(output)
	if absIn == absOut {
		return fmt.Errorf("input and output paths must be different")
	}

	return nil
}
