package config

import (
	"github.com/Garmonik/gRPC_chat/backend/auth/internal/pkg/config_lib"
	"github.com/lpernett/godotenv"
	"time"
)

type DBConfig struct {
	Name     string `env:"DB_NAME"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	Host     string `env:"DB_HOST"`
	Port     int    `env:"DB_PORT"`
	SSLMode  string `env:"DB_SSLMODE"`
}

type GRPCConfig struct {
	Port    int           `env:"GRPC_PORT"`
	Timeout time.Duration `env:"GRPC_TIMEOUT"`
}

type Config struct {
	DB            DBConfig      `envPrefix:"DB_"`
	GRPC          GRPCConfig    `envPrefix:"GRPC_"`
	Environment   string        `env:"ENVIRONMENT"`
	SessionTTL    time.Duration `env:"SESSION_TTL"`
	Argon2Time    uint8         `env:"ARGON2_TIME"`
	Argon2Memory  uint8         `env:"ARGON2_MEMORY"`
	Argon2Threads uint8         `env:"ARGON2_THREADS"`
	Argon2KeyLen  uint8         `env:"ARGON2_KEY_LEN"`
	Secret        string        `env:"SECRET"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		DB:            loadDBConfig(),
		GRPC:          loadGRPCConfig(),
		Environment:   config_lib.GetEnvStr("ENVIRONMENT"),
		SessionTTL:    config_lib.GetEnvDuration("SESSION_TTL"),
		Argon2Time:    uint8(config_lib.GetEnvInt("ARGON2_TIME")),
		Argon2Memory:  uint8(config_lib.GetEnvInt("ARGON2_MEMORY")),
		Argon2Threads: uint8(config_lib.GetEnvInt("ARGON2_THREADS")),
		Argon2KeyLen:  uint8(config_lib.GetEnvInt("ARGON2_KEY_LEN")),
		Secret:        config_lib.GetEnvStr("SECRET"),
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
		SSLMode:  config_lib.GetEnvStr("DB_SSLMODE"),
	}
}
