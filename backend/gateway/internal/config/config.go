package config

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/pkg/config_lib"
	"github.com/lpernett/godotenv"
)

type DBConfig struct {
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	SSLMode  string `env:"DB_SSLMODE"`
}

type Config struct {
	DB          DBConfig `envPrefix:"DB_"`
	Environment string   `env:"ENVIRONMENT"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		DB:          loadDBConfig(),
		Environment: config_lib.GetEnvStr("ENVIRONMENT"),
	}
}

func loadDBConfig() DBConfig {
	return DBConfig{
		Name:     config_lib.GetEnvStr("DB_NAME"),
		User:     config_lib.GetEnvStr("DB_USER"),
		Password: config_lib.GetEnvStr("DB_PASSWORD"),
		Host:     config_lib.GetEnvStr("DB_HOST"),
		Port:     config_lib.GetEnvInt("DB_PORT"),
		SSLMode:  config_lib.GetEnvStr("DB_SSLMODE"),
	}
}
