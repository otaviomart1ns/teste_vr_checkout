package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/streadway/amqp"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/config"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/infra/db/sqlc"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/infra/gateway"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/api/handlers"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/queue"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/usecases"

	_ "github.com/lib/pq"
)

// @title API - VR Checkout
// @version 1.0
// @description Documentação da API REST para o teste da VR Software
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Carrega variáveis de ambiente
	cfg := config.Load()

	// Conecta ao PostgreSQL
	pool, err := pgxpool.New(context.Background(), cfg.PostgresURL)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados:", err)
	}
	defer pool.Close()

	// Conecta ao RabbitMQ
	rabbitConn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		log.Fatal("Erro ao conectar ao RabbitMQ:", err)
	}
	defer rabbitConn.Close()

	// Repositório SQLC
	repo := sqlc.NewTransactionRepository(pool)

	// Clients externos
	treasuryClient := gateway.NewTreasuryClient(cfg)
	currencyMetaClient := gateway.NewCurrencyMetaClient(cfg)

	// Producer e Consumer
	producer, err := queue.NewTransactionProducer(rabbitConn)
	if err != nil {
		log.Fatal("Erro ao criar producer:", err)
	}

	consumer, err := queue.NewTransactionConsumer(rabbitConn, repo)
	if err != nil {
		log.Fatal("Erro ao criar consumer:", err)
	}
	if err := consumer.StartConsuming(); err != nil {
		log.Fatal("Erro ao iniciar consumo da fila:", err)
	}

	// Caso de uso
	service := usecases.NewTransactionService(producer, repo, treasuryClient)

	// Handlers
	transactionHandler := handlers.NewTransactionHandler(service)
	currencyHandler := handlers.NewCurrencyHandler(currencyMetaClient)

	// API
	router := api.SetupRouter(transactionHandler, currencyHandler)

	log.Println("Servidor iniciado em http://localhost:" + cfg.ServerPort)
	log.Fatal(router.Run(":" + cfg.ServerPort))
}
