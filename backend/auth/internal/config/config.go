package config

import (
	"github.com/Garmonik/gRPC_chat/backend/auth/pkg/config_lib"
	"github.com/lpernett/godotenv"
	"time"
)

type DBConfig struct {
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
}

type GRPCConfig struct {
	Port    int           `env:"GRPC_PORT"`
	Timeout time.Duration `env:"GRPC_TIMEOUT"`
}

type Config struct {
	DB          DBConfig      `envPrefix:"DB_"`
	GRPC        GRPCConfig    `envPrefix:"GRPC_"`
	Environment string        `env:"ENVIRONMENT"`
	TokenTTL    time.Duration `env:"TOKEN_TTL"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		DB:          loadDBConfig(),
		GRPC:        loadGRPCConfig(),
		Environment: config_lib.GetEnvStr("ENVIRONMENT"),
		TokenTTL:    config_lib.GetEnvDuration("TOKEN_TTL"),
	}
}

func loadGRPCConfig() GRPCConfig {
	return GRPCConfig{
		Port:    config_lib.GetEnvInt("GRPC_PORT"),
		Timeout: config_lib.GetEnvDuration("GRPC_TIMEOUT"),
	}
}

func loadDBConfig() DBConfig {
	return DBConfig{
		Name:     config_lib.GetEnvStr("DB_NAME"),
		User:     config_lib.GetEnvStr("DB_USER"),
		Password: config_lib.GetEnvStr("DB_PASSWORD"),
		Host:     config_lib.GetEnvStr("DB_HOST"),
		Port:     config_lib.GetEnvInt("DB_PORT"),
	}
}
