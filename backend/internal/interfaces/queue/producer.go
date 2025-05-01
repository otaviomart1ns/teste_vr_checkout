package queue

import (
	"context"
	"encoding/json"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/streadway/amqp"
)

type TransactionProducer struct {
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewTransactionProducer(conn *amqp.Connection) (*TransactionProducer, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return &TransactionProducer{
		Channel: ch,
		Queue:   q,
	}, nil
}

func (p *TransactionProducer) PublishTransaction(_ context.Context, tx *entities.Transaction) error {
	body, err := json.Marshal(tx)
	if err != nil {
		return err
	}

	return p.Channel.Publish(
		"",           // exchange
		p.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}
