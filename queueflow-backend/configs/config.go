package config

import (
	"os"
	"strconv"
)

type Config struct {
	RedisURL    string
	WorkerCount int
}

func Load() Config {

	workers, err := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	if err != nil {
		workers = 8
	}

	redisURL := getEnv("REDIS_URL", "localhost:6379")

	return Config{
		RedisURL:    redisURL,
		WorkerCount: workers,
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
