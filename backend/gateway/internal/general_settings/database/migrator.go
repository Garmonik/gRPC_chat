package database

import (
	"database/sql"
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/general_settings/config"
	_ "github.com/lib/pq"
	"github.com/urfave/cli/v2"
	"log/slog"

	"github.com/pressly/goose/v3"
)

type MigrationFunc func(db *sql.DB, dir string, opts ...goose.OptionsFunc) error

func CreateMigratorApp(cfg *config.Config, log *slog.Logger) *cli.App {
	return &cli.App{
		Name:  "migrator",
		Usage: "Database migration tool",
		Commands: []*cli.Command{
			createUpCommand(cfg, log),
			createDownCommand(cfg, log),
			createStatusCommand(cfg, log),
			createVersionCommand(cfg, log),
		},
	}
}

func createUpCommand(cfg *config.Config, log *slog.Logger) *cli.Command {
	return &cli.Command{
		Name:  "up",
		Usage: "Migrate the DB to the most recent version available",
		Action: func(c *cli.Context) error {
			return RunMigrations(cfg, log, goose.Up)
		},
	}
}

func createDownCommand(cfg *config.Config, log *slog.Logger) *cli.Command {
	return &cli.Command{
		Name:  "down",
		Usage: "Roll back the version by 1",
		Action: func(c *cli.Context) error {
			return RunMigrations(cfg, log, goose.Down)
		},
	}
}

func createStatusCommand(cfg *config.Config, log *slog.Logger) *cli.Command {
	return &cli.Command{
		Name:  "status",
		Usage: "Dump the migration status for the current DB",
		Action: func(c *cli.Context) error {
			return RunMigrations(cfg, log, func(db *sql.DB, dir string, opts ...goose.OptionsFunc) error {
				return goose.Status(db, dir)
			})
		},
	}
}

func createVersionCommand(cfg *config.Config, log *slog.Logger) *cli.Command {
	return &cli.Command{
		Name:  "version",
		Usage: "Print the current version of the database",
		Action: func(c *cli.Context) error {
			return RunMigrations(cfg, log, func(db *sql.DB, dir string, opts ...goose.OptionsFunc) error {
				version, err := goose.GetDBVersion(db)
				if err != nil {
					return err
				}
				log.Info("current database version", "version", version)
				return nil
			})
		},
	}
}

func RunMigrations(cfg *config.Config, log *slog.Logger, op MigrationFunc) error {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Name,
		cfg.DB.SSLMode,
	)

	log.Info("connecting to database",
		"host", cfg.DB.Host,
		"name", cfg.DB.Name,
		"user", cfg.DB.User,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("failed to open DB connection: %w", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping DB: %w", err)
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	migrationsDir := "./migrations"
	log.Info("executing migration operation", "dir", migrationsDir)

	if err := op(db, migrationsDir); err != nil {
		return fmt.Errorf("migration operation failed: %w", err)
	}

	return nil
}
