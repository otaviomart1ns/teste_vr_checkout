package queue_test

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/config"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/domain/entities"
	"github.com/otaviomart1ns/teste_vr_checkout/backend/internal/interfaces/queue"
	"github.com/streadway/amqp"
	"github.com/stretchr/testify/assert"
)

func setupTestProducer(t *testing.T) (*queue.TransactionProducer, *amqp.Channel, amqp.Queue) {
	t.Helper()

	cfg := config.Load()

	conn, err := amqp.Dial(cfg.RabbitMQURL)
	if err != nil {
		t.Fatalf("erro ao conectar no RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		t.Fatalf("erro ao abrir canal: %v", err)
	}

	queueName := "transactions_test"
	q, err := ch.QueueDeclare(
		queueName,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
	if err != nil {
		t.Fatalf("erro ao declarar fila: %v", err)
	}

	producer := &queue.TransactionProducer{
		Channel: ch,
		Queue:   q,
	}

	return producer, ch, q
}

func TestPublishTransaction_Integration(t *testing.T) {
	producer, ch, q := setupTestProducer(t)
	defer ch.Close()

	tx := &entities.Transaction{
		ID:          "test-123",
		Description: "Transação de Teste",
		Date:        time.Now().UTC(),
		ValueUSD:    99.99,
	}

	err := producer.PublishTransaction(context.Background(), tx)
	assert.NoError(t, err)

	msgs, err := ch.Consume(q.Name, "", true, true, false, false, nil)
	assert.NoError(t, err)

	select {
	case msg := <-msgs:
		var received entities.Transaction
		err := json.Unmarshal(msg.Body, &received)
		assert.NoError(t, err)
		assert.Equal(t, tx.ID, received.ID)
		assert.Equal(t, tx.Description, received.Description)
		assert.InDelta(t, tx.ValueUSD, received.ValueUSD, 0.01)
	case <-time.After(2 * time.Second):
		t.Fatal("timeout esperando mensagem da fila")
	}
}

func TestNewTransactionProducer(t *testing.T) {
	cfg := config.Load()

	conn, err := amqp.Dial(cfg.RabbitMQURL)
	assert.NoError(t, err)
	defer conn.Close()

	producer, err := queue.NewTransactionProducer(conn)
	assert.NoError(t, err)
	assert.NotNil(t, producer)
	assert.NotNil(t, producer.Channel)
	assert.NotEmpty(t, producer.Queue.Name)
}

func TestNewTransactionProducer_ChannelError(t *testing.T) {
	cfg := config.Load()
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	assert.NoError(t, err)
	conn.Close()

	_, err = queue.NewTransactionProducer(conn)
	assert.Error(t, err)
}
