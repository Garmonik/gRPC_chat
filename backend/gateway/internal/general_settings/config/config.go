package config

import (
	"github.com/Garmonik/gRPC_chat/backend/gateway/internal/pkg/config_lib"
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

type GRPCAuthConfig struct {
	Port int    `env:"GRPC_AUTH_PORT"`
	Host string `env:"GRPC_AUTH_HOST"`
}

type HTTPServerConfig struct {
	Address     string
	Timeout     time.Duration `env:"HTTP_AUTH_TIMEOUT" env-default:"4s"`
	IdleTimeout time.Duration `env:"HTTP_AUTH_IDLE_TIMEOUT" env-default:"60s"`
}

type Config struct {
	DB          DBConfig         `envPrefix:"DB_"`
	Environment string           `env:"ENVIRONMENT"`
	GrpcAuth    GRPCAuthConfig   `envPrefix:"GRPC_AUTH_"`
	HTTPServer  HTTPServerConfig `envPrefix:"HTTP_"`
}

func MustLoad() *Config {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	return &Config{
		DB:          loadDBConfig(),
		Environment: config_lib.GetEnvStr("ENVIRONMENT"),
		GrpcAuth:    loadGrpcAuthConfig(),
		HTTPServer:  loadHTTPServerConfig(),
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

func loadGrpcAuthConfig() GRPCAuthConfig {
	return GRPCAuthConfig{
		Host: config_lib.GetEnvStr("GRPC_AUTH_HOST"),
		Port: config_lib.GetEnvInt("GRPC_AUTH_PORT"),
	}
}

func loadHTTPServerConfig() HTTPServerConfig {
	return HTTPServerConfig{
		Address:     config_lib.GetEnvAddress("HTTP_HOST", "HTTP_PORT"),
		Timeout:     config_lib.GetEnvDuration("HTTP_TIMEOUT"),
		IdleTimeout: config_lib.GetEnvDuration("HTTP_IDLE_TIMEOUT"),
	}
}
