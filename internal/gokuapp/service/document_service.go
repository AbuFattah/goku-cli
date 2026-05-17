package service

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/abufattah/goku-cli/internal/cli/root"
	"github.com/abufattah/goku-cli/internal/cli/save"
	"github.com/abufattah/goku-cli/internal/gokuapp/repository"
	database "github.com/abufattah/goku-cli/internal/platform/database/sqlc"
	"github.com/abufattah/goku-cli/internal/platform/fileutil"
)

type DocumentService struct {
	documentRepo *repository.DocumentRepository
}

func NewDocumentService(ctx context.Context, docRepo *repository.DocumentRepository) save.DocumentService {
	return &DocumentService{
		documentRepo: docRepo,
	}
}

func (ds *DocumentService) SaveDocument(ctx context.Context, path string) error {
	base := filepath.Base(path)
	ext := filepath.Ext(path)

	name := strings.TrimSuffix(base, ext)
	format, err := fileutil.DetectFormat(path)
	if err != nil {
		return fmt.Errorf("unsupported file format")
	}

	data, err := fileutil.Read(path)
	if err != nil {
		return fmt.Errorf("unsupported data")
	}

	if format == fileutil.FormatYAML {
		data, err = root.Convert(data, fileutil.FormatYAML, fileutil.FormatJSON)
		if err != nil {
			return fmt.Errorf("invalid file format")
		}
	}

	_, err = ds.documentRepo.SaveDocument(ctx, database.SaveDocumentParams{
		Name:       name,
		DataFormat: string(format),
		Data:       data,
	})

	if err != nil {
		return fmt.Errorf("document save failed")
	}

	return nil
}
