package config

import (
	"os"
	"strconv"
)

type Config struct {
	RedisURL    string
	WorkerCount int

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func Load() Config {

	workers, err := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	if err != nil {
		workers = 8
	}

	redisURL := getEnv("REDIS_URL", "localhost:6379")

	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbName := getEnv("DB_NAME", "queueflow")

	return Config{
		RedisURL:    redisURL,
		WorkerCount: workers,
		DBHost:      dbHost,
		DBPort:      dbPort,
		DBUser:      dbUser,
		DBPassword:  dbPassword,
		DBName:      dbName,
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
