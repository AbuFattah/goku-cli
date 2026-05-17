package container

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/abufattah/goku-cli/config"
	"github.com/abufattah/goku-cli/internal/cli/root"
	"github.com/abufattah/goku-cli/internal/cli/save"

	"github.com/abufattah/goku-cli/internal/gokuapp/repository"
	"github.com/abufattah/goku-cli/internal/gokuapp/service"
	database "github.com/abufattah/goku-cli/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5"
)

type Container struct {
	cfg             *config.Config
	postgres        *pgx.Conn
	Queries         *database.Queries
	logger          *slog.Logger
	documentService save.DocumentService
}

func New() *Container {
	return &Container{}
}

func (c *Container) Run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("loading config: %w", err)
	}
	c.cfg = cfg

	c.logger = newLogger(cfg.LogLevel)
	c.logger.Info("starting application", "env", cfg.AppEnv)

	rootCmd := root.NewRootCommand(c.initializePostgres)
	rootCmd.AddCommand(save.NewSaveCommand(ctx, c.documentServiceProvider))

	return rootCmd.ExecuteContext(ctx)
}

func (c *Container) Close() {
	if c.logger != nil {
		c.logger.Info("shutting down application")
	}
	if c.postgres != nil {
		c.postgres.Close(context.Background())
	}
}

func newLogger(levelStr string) *slog.Logger {
	var level slog.Level
	if err := level.UnmarshalText([]byte(levelStr)); err != nil {
		level = slog.LevelInfo
	}
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
}

func (c *Container) documentServiceProvider(ctx context.Context) (save.DocumentService, error) {

	if c.documentService != nil {
		return c.documentService, nil
	}

	if err := c.initializePostgres(ctx); err != nil {
		return nil, fmt.Errorf("connecting to database: %w", err)
	}

	docRepo := repository.NewDocumentRepository(c.Queries)
	c.documentService = service.NewDocumentService(ctx, docRepo)

	return c.documentService, nil
}
