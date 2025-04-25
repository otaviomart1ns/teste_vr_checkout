package config

import (
	"log"
	"os"
)

type Config struct {
	PostgresURL string
	RabbitMQURL string
}

func Load() *Config {
	postgresURL, ok := os.LookupEnv("POSTGRES_URL")
	if !ok || postgresURL == "" {
		log.Fatal("POSTGRES_URL environment variable is required")
	}

	rabbitMQURL, ok := os.LookupEnv("RABBITMQ_URL")
	if !ok || rabbitMQURL == "" {
		log.Fatal("RABBITMQ_URL environment variable is required")
	}

	return &Config{
		PostgresURL: postgresURL,
		RabbitMQURL: rabbitMQURL,
	}
}