package config

import (
	postgresconn "github.com/IrentyM/store/statistics-service/pkg/postgres"
	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Postgres postgresconn.PostgresConfig
		NATS     NATSConfig
		Server   Server
	}

	Server struct {
		GRPCServer GRPCServer
	}

	GRPCServer struct {
		Port int    `env:"GRPC_PORT,required"`
		Mode string `env:"GIN_MODE" envDefault:"release"`
	}

	NATSConfig struct {
		URL string `env:"NATS_URL,required"`
	}
)

func New() (*Config, error) {
	var cfg Config
	opts := env.Options{Environment: nil, TagName: "env", Prefix: ""}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.Parse(&cfg, opts)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
