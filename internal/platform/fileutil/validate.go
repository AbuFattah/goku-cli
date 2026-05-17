package fileutil

import (
	"fmt"
	"os"
)

func ValidateFilepathFlags(input, output string) error {

	if input == "" {
		return fmt.Errorf("input file is required (-i/--input)")
	}
	if output == "" {
		return fmt.Errorf("output file is required (-o/--output)")
	}

	inFormat, err := ValidateFile(input)
	if err != nil {
		return err
	}

	outFormat, err := ValidateFile(output)
	if err != nil {
		return err
	}

	if inFormat == outFormat {
		return fmt.Errorf("input and output formats should be different")
	}

	return nil
}

func ValidateFile(input string) (Format, error) {
	file, err := os.Stat(input)

	if os.IsNotExist(err) {
		return "", fmt.Errorf("input file does not exist: %s", input)
	}

	if file.IsDir() {
		return "", fmt.Errorf("the file path is a directory: %s", input)
	}

	format, err := DetectFormat(input)

	if err != nil {
		return "", fmt.Errorf("unsupported file format")
	}

	return format, nil
}
