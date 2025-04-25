package queue

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/usecases"
)

type TransactionMessage struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Date        string    `json:"date"`
	Amount      float64   `json:"amount"`
}

type TransactionConsumer struct {
	service   usecases.TransactionService
	channel   *amqp.Channel
	queueName string
}

func NewTransactionConsumer(service usecases.TransactionService, ch *amqp.Channel, queueName string) *TransactionConsumer {
	return &TransactionConsumer{
		service:   service,
		channel:   ch,
		queueName: queueName,
	}
}

func (c *TransactionConsumer) StartConsuming() error {
	msgs, err := c.channel.Consume(
		c.queueName,
		"",
		true,  // auto-ack
		false, // exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			var txMsg TransactionMessage
			if err := json.Unmarshal(msg.Body, &txMsg); err != nil {
				log.Printf("failed to unmarshal message: %v", err)
				continue
			}

			date, err := time.Parse("2006-01-02", txMsg.Date)
			if err != nil {
				log.Printf("invalid date in message: %v", err)
				continue
			}

			_, err = c.service.CreateTransaction(context.Background(), txMsg.Description, date, txMsg.Amount)
			if err != nil {
				log.Printf("failed to save transaction: %v", err)
				continue
			}

			log.Printf("transaction %s persisted successfully", txMsg.ID)
		}
	}()

	return nil
}
