package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/config"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/infra/db"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/infra/queue"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/usecases"
)

func main() {
	cfg := config.Load()

	// PostgreSQL
	db.Connect(cfg.PostgresURL)
	queries := db.Queries

	// Service
	service := usecases.NewTransactionService(queries)

	// RabbitMQ
	rmq, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer rmq.Close()

	channel, err := rmq.Channel()
	if err != nil {
		log.Fatalf("Failed to open RabbitMQ channel: %v", err)
	}

	// Start consumer
	consumer := queue.NewTransactionConsumer(service, channel, "transactions.create")
	if err := consumer.StartConsuming(); err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}

	// Start publisher and routes
	publisher, _ := queue.NewRabbitMQPublisher(rmq, "transactions.create")
	router := gin.Default()
	api.RegisterTransactionRoutes(router, publisher)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}