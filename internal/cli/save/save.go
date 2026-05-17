package save

import (
	"context"
	"fmt"

	"github.com/abufattah/goku-cli/internal/platform/fileutil"
	"github.com/spf13/cobra"
)

type DocumentService interface {
	SaveDocument(ctx context.Context, path string) error
}

type ServiceProvider func(ctx context.Context) (DocumentService, error)

func NewSaveCommand(ctx context.Context, docService ServiceProvider) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "save",
		Short: "Dump data to documents table",
		Long:  "Dump data to documents table",
		RunE: func(cmd *cobra.Command, args []string) error {
			input, _ := cmd.Flags().GetString("input")

			if input == "" {
				return cmd.Help()
			}

			if _, err := fileutil.ValidateFile(input); err != nil {
				return err
			}

			ds, err := docService(ctx)
			if err != nil {
				return err
			}

			if err := ds.SaveDocument(ctx, input); err != nil {
				return fmt.Errorf("failed to save document")
			}

			return nil
		},
	}

	cmd.Flags().StringP("input", "i", "", "Path to input file")

	return cmd
}
