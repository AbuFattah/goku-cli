package root

import (
	"context"
	"fmt"

	"github.com/abufattah/goku-cli/config"
	"github.com/abufattah/goku-cli/internal/platform/fileutil"
	"github.com/spf13/cobra"
)

func NewRootCommand(cfg *config.Config, initPostgres func(ctx context.Context) error) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goku",
		Short: "A Go CLI for file conversion",
		Long:  `A Go CLI for file conversion`,
		RunE: func(cmd *cobra.Command, args []string) error {
			input, _ := cmd.Flags().GetString("input")
			output, _ := cmd.Flags().GetString("output")

			if input == "" && output == "" {
				return cmd.Help()
			}

			if err := validateFilepathFlags(input, output); err != nil {
				return err
			}

			return ExecuteConversion(input, output)
		},
	}

	cmd.Flags().StringP("input", "i", "", "Path to input file")
	cmd.Flags().StringP("output", "o", "", "Path to output file")

	return cmd
}

func ExecuteConversion(input, output string) error {
	inputData, err := fileutil.Read(input)
	if err != nil {
		return err
	}

	inputFormat, err := DetectFormat(input)
	if err != nil {
		return fmt.Errorf("input: %w", err)
	}

	outFormat, err := DetectFormat(output)
	if err != nil {
		return fmt.Errorf("output: %w", err)
	}

	outData, err := Convert(inputData, inputFormat, outFormat)
	if err != nil {
		return fmt.Errorf("converting %s → %s: %w", inputFormat, outFormat, err)
	}

	return fileutil.Write(output, outData)
}
