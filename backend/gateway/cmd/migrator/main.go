package main

import (
	"database/sql"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/config"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/database"
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/pkg/logger"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Environment)

	app := &cli.App{
		Name:  "migrator",
		Usage: "Database migration tool",
		Commands: []*cli.Command{
			{
				Name:  "up",
				Usage: "Migrate the DB to the most recent version available",
				Action: func(c *cli.Context) error {
					return database.RunMigrations(cfg, log, goose.Up)
				},
			},
			{
				Name:  "down",
				Usage: "Roll back the version by 1",
				Action: func(c *cli.Context) error {
					return database.RunMigrations(cfg, log, goose.Down)
				},
			},
			{
				Name:  "status",
				Usage: "Dump the migration status for the current DB",
				Action: func(c *cli.Context) error {
					return database.RunMigrations(cfg, log, func(db *sql.DB, dir string, opts ...goose.OptionsFunc) error {
						return goose.Status(db, dir)
					})
				},
			},
			{
				Name:  "version",
				Usage: "Print the current version of the database",
				Action: func(c *cli.Context) error {
					return database.RunMigrations(cfg, log, func(db *sql.DB, dir string, opts ...goose.OptionsFunc) error {
						version, err := goose.GetDBVersion(db)
						if err != nil {
							return err
						}
						log.Info("current database version", "version", version)
						return nil
					})
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Error("migration failed", "error", err)
	}
}
