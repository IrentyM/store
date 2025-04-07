package server

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBhost     string
	DBport     string
	DBuser     string
	DBpassword string
	DBname     string
}

var ErrInvalidEnv = errors.New("invalid environment variables")

func GetConfig() *Config {
	config, err := ParseEnvConfig()
	if err != nil {
		return GetDefaultConfig()
	}
	return config
}

func GetDefaultConfig() *Config {
	return &Config{
		Port:       ":8080",
		DBhost:     "localhost",
		DBport:     "5432",
		DBuser:     "postgres",
		DBpassword: "postgres",
		DBname:     "orders",
	}
}

func ParseEnvConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	var (
		port       = os.Getenv("PORT")
		dbHost     = os.Getenv("DB_HOST")
		dbPort     = os.Getenv("DB_PORT")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
	)

	if port == "" || dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return nil, ErrInvalidEnv
	}

	return &Config{
		Port:       ":" + port,
		DBhost:     dbHost,
		DBport:     dbPort,
		DBuser:     dbUser,
		DBpassword: dbPassword,
		DBname:     dbName,
	}, nil
}
