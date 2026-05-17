package container

import (
	"context"
	"fmt"
	"log/slog"

	database "github.com/abufattah/goku-cli/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5"
)

func (c *Container) initializePostgres(ctx context.Context) error {
	c.logger.Info("initializing postgres connection pool")

	pg, err := pgx.Connect(ctx, c.cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect to postgres database", "error", err)
		return fmt.Errorf("failed to connect to postgres database: %w", err)
	}

	if err := pg.Ping(ctx); err != nil {
		slog.Error("failed to ping postgres database", "error", err)
		pg.Close(ctx)
		return fmt.Errorf("failed to ping postgres database: %w", err)
	}

	c.postgres = pg
	c.Queries = database.New(pg)

	c.logger.Info("Database connected successfully")

	return nil
}
