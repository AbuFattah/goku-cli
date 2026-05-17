package container

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	database "github.com/abufattah/goku-cli/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (c *Container) initializePostgres(ctx context.Context) error {
	slog.Info("initializing postgres connection pool")

	config, err := pgxpool.ParseConfig(c.cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to parse database configuration", "error", err)
		return fmt.Errorf("failed to parse database configuration: %w", err)
	}

	config.MaxConns = 30
	config.MinConns = 15
	config.MaxConnLifetime = 20 * time.Minute
	config.MaxConnIdleTime = 10 * time.Minute

	pgPool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		slog.Error("failed to create postgres connection pool", "error", err)
		return fmt.Errorf("failed to create postgres connection pool: %w", err)
	}

	if err := pgPool.Ping(ctx); err != nil {
		slog.Error("failed to ping postgres database", "error", err)
		pgPool.Close()
		return fmt.Errorf("failed to ping postgres database: %w", err)
	}

	c.postgres = pgPool
	c.Queries = database.New(pgPool)

	c.logger.Info("Database connection pool configured successfully",
		slog.Int("maxConns", int(config.MaxConns)),
		slog.Int("minConns", int(config.MinConns)),
		slog.Duration("maxConnLifetime", config.MaxConnLifetime),
		slog.Duration("maxConnIdleTime", config.MaxConnIdleTime),
	)

	return nil
}
