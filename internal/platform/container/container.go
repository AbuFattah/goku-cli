package container

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/abufattah/goku-cli/config"
	"github.com/abufattah/goku-cli/internal/cli/root"
	database "github.com/abufattah/goku-cli/internal/platform/database/sqlc"
	"github.com/jackc/pgx/v5"
)

type Container struct {
	cfg      *config.Config
	postgres *pgx.Conn
	Queries  *database.Queries
	logger   *slog.Logger
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

	rootCmd := root.NewRootCommand(cfg, c.initializePostgres)

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
