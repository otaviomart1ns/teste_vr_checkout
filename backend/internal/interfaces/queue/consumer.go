package queue

import (
	"context"
	"encoding/json"
	"log"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/streadway/amqp"
)

type TransactionSaver interface {
	Save(ctx context.Context, tx *entities.Transaction) (string, error)
}

type TransactionConsumer struct {
	conn *amqp.Connection
	repo interface {
		Save(ctx context.Context, tx *entities.Transaction) error
	}
}

func NewTransactionConsumer(conn *amqp.Connection, repo interface {
    Save(ctx context.Context, tx *entities.Transaction) error
}) (*TransactionConsumer, error) {
	return &TransactionConsumer{
		conn: conn,
		repo: repo,
	}, nil
}

func (c *TransactionConsumer) StartConsuming() error {
	ch, err := c.conn.Channel()
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		"transactions",
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name,
		"",    // consumer tag
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			var tx entities.Transaction
			if err := json.Unmarshal(d.Body, &tx); err != nil {
				log.Println("Erro ao decodificar mensagem:", err)
				continue
			}

			if err := c.repo.Save(context.Background(), &tx); err != nil {
				log.Println("Erro ao salvar no banco:", err)
				continue
			}

			log.Printf("Transação salva: %s\n", tx.ID)
		}
	}()

	return nil
}