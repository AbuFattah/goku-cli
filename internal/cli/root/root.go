package root

import (
	"github.com/abufattah/goku-cli/config"
	"github.com/spf13/cobra"
)

func NewRootCommand(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "goku",
		Short: "A Go CLI for file conversion",
		Long:  `A Go CLI for file conversion`,
	}

	cmd.Flags().StringP("input", "i", "", "Path to input file")
	cmd.Flags().StringP("output", "o", "", "Path to output file")

	return cmd
}
