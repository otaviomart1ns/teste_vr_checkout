package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


	type Config struct {
		PostgresURL string
		RabbitMQURL string
		TreasuryBaseURL string
		TreasureEndpont string
		ServerPort string
	}

	func Load() *Config {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Println("⚠️ Warning: .env não encontrado no caminho esperado")
		}
		
		postgresURL, ok := os.LookupEnv("POSTGRES_URL")
		if !ok || postgresURL == "" {
			log.Fatal("POSTGRES_URL environment variable is required")
		}

		rabbitMQURL, ok := os.LookupEnv("RABBITMQ_URL")
		if !ok || rabbitMQURL == "" {
			log.Fatal("RABBITMQ_URL environment variable is required")
		}

		treasuryBaseURL, ok := os.LookupEnv("TREASURY_API_BASE_URL")
		if !ok || treasuryBaseURL == "" {
			log.Fatal("TREASURY_API_BASE_URL environment variable is required")
		}

		treasuryEndpoint, ok := os.LookupEnv("TREASURY_API_ENDPOINT")
		if !ok || treasuryEndpoint == "" {
			log.Fatal("TREASURY_API_ENDPOINT environment variable is required")
		}

		serverPort, ok := os.LookupEnv("API_PORT")
		if !ok || serverPort == "" {
			log.Fatal("API_PORT environment variable is required")
		}

		return &Config{
			PostgresURL: postgresURL,
			RabbitMQURL: rabbitMQURL,
			TreasuryBaseURL: treasuryBaseURL,
			TreasureEndpont: treasuryEndpoint,
			ServerPort: serverPort,
		}
	}