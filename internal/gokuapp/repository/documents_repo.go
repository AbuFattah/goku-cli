package repository

import (
	"context"

	database "github.com/abufattah/goku-cli/internal/platform/database/sqlc"
)

type DocumentRepository struct {
	queries *database.Queries
}

func NewDocumentRepository(queries *database.Queries) *DocumentRepository {
	return &DocumentRepository{queries: queries}
}

func (dr *DocumentRepository) SaveDocument(ctx context.Context, arg database.SaveDocumentParams) (database.Document, error) {
	return dr.queries.SaveDocument(ctx, arg)
}
