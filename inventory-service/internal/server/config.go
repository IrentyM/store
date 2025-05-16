package server

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port       string
	DBhost     string
	DBport     string
	DBuser     string
	DBpassword string
	DBname     string
	RedisAddr  string
	RedisPass  string
	RedisDB    int
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
		Port:       ":8070",
		DBhost:     "localhost",
		DBport:     "5432",
		DBuser:     "postgres",
		DBpassword: "postgres",
		DBname:     "store",
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
		redisAddr  = os.Getenv("REDIS_ADDR")
		redisPass  = os.Getenv("REDIS_PASS")
		redisDB    = os.Getenv("REDIS_DB")
	)

	if port == "" || dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return nil, ErrInvalidEnv
	}

	// Default Redis values if not set
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	if redisDB == "" {
		redisDB = "0"
	}

	dbNum, _ := strconv.Atoi(redisDB)

	return &Config{
		Port:       ":" + port,
		DBhost:     dbHost,
		DBport:     dbPort,
		DBuser:     dbUser,
		DBpassword: dbPassword,
		DBname:     dbName,
		RedisAddr:  redisAddr,
		RedisPass:  redisPass,
		RedisDB:    dbNum,
	}, nil
}
