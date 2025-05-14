package database

import (
	"fmt"
	"github.com/Garmonik/gRPC_chat/backend/user/internal/general_settings/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

type DataBase struct {
	Db *gorm.DB
}

func New(log *slog.Logger, cfg *config.Config) (*DataBase, error) {
	const op = "auth.database.New"
	log = log.With(slog.String("op", op))
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.DB.Host,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.Name,
		cfg.DB.Port,
		cfg.DB.SSLMode,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("failed to connect to database", slog.String("error", err.Error()))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	log.Info("database connection established")
	return &DataBase{Db: database}, nil
}

func (db *DataBase) Close() error {
	if db.Db == nil {
		return nil
	}

	sqlDB, err := db.Db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB: %w", err)
	}
	return sqlDB.Close()
}
