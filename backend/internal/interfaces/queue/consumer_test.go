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
	"github.com/stretchr/testify/mock"
)

type MockRepo struct {
	mock.Mock
}

func (m *MockRepo) Save(ctx context.Context, tx *entities.Transaction) error {
	args := m.Called(ctx, tx)
	return args.Error(0)
}

func TestTransactionConsumer_Integration(t *testing.T) {
	cfg := config.Load()

	conn, err := amqp.Dial(cfg.RabbitMQURL)
	assert.NoError(t, err)
	defer conn.Close()

	ch, err := conn.Channel()
	assert.NoError(t, err)
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"transactions",
		true,
		false,
		false,
		false,
		nil,
	)
	assert.NoError(t, err)

	mockRepo := new(MockRepo)

	consumer, err := queue.NewTransactionConsumer(conn, mockRepo)
	assert.NoError(t, err)

	done := make(chan bool)

	tx := &entities.Transaction{
		ID:          "consumer-test-id",
		Description: "Teste Consumer",
		Date:        time.Now().UTC(),
		ValueUSD:    55.55,
	}

	mockRepo.On("Save", mock.Anything, mock.MatchedBy(func(received *entities.Transaction) bool {
		return received.ID == tx.ID &&
			received.Description == tx.Description &&
			received.ValueUSD == tx.ValueUSD
	})).Return(nil).Run(func(args mock.Arguments) {
		done <- true
	})

	// Inicia o consumo
	err = consumer.StartConsuming()
	assert.NoError(t, err)

	// Publica a transação na fila
	body, err := json.Marshal(tx)
	assert.NoError(t, err)

	err = ch.Publish(
		"",           // exchange
		q.Name,       // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	assert.NoError(t, err)

	select {
	case <-done:
		mockRepo.AssertCalled(t, "Save", mock.Anything, mock.AnythingOfType("*entities.Transaction"))
	case <-time.After(3 * time.Second):
		t.Fatal("timeout: transação não foi processada")
	}
}

func TestTransactionConsumer_UnmarshalError(t *testing.T) {
	cfg := config.Load()

	conn, err := amqp.Dial(cfg.RabbitMQURL)
	assert.NoError(t, err)
	defer conn.Close()

	mockRepo := new(MockRepo)

	consumer, err := queue.NewTransactionConsumer(conn, mockRepo)
	assert.NoError(t, err)

	err = consumer.StartConsuming()
	assert.NoError(t, err)

	ch, err := conn.Channel()
	assert.NoError(t, err)
	defer ch.Close()

	err = ch.Publish(
		"",
		"transactions",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("isso_nao_e_json"),
		},
	)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	mockRepo.AssertNotCalled(t, "Save", mock.Anything, mock.Anything)
}

func TestTransactionConsumer_SaveError(t *testing.T) {
	cfg := config.Load()

	conn, err := amqp.Dial(cfg.RabbitMQURL)
	assert.NoError(t, err)
	defer conn.Close()

	ch, err := conn.Channel()
	assert.NoError(t, err)
	defer ch.Close()

	tx := &entities.Transaction{
		ID:          "fail-save-id",
		Description: "Teste Save Error",
		Date:        time.Now().UTC(),
		ValueUSD:    77.77,
	}

	body, err := json.Marshal(tx)
	assert.NoError(t, err)

	mockRepo := new(MockRepo)
	mockRepo.On("Save", mock.Anything, mock.AnythingOfType("*entities.Transaction")).
		Return(assert.AnError)

	consumer, err := queue.NewTransactionConsumer(conn, mockRepo)
	assert.NoError(t, err)

	err = consumer.StartConsuming()
	assert.NoError(t, err)

	err = ch.Publish(
		"",
		"transactions",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	assert.NoError(t, err)

	time.Sleep(1 * time.Second)

	mockRepo.AssertCalled(t, "Save", mock.Anything, mock.AnythingOfType("*entities.Transaction"))
}

func TestStartConsuming_ChannelError(t *testing.T) {
	cfg := config.Load()
	conn, err := amqp.Dial(cfg.RabbitMQURL)
	assert.NoError(t, err)
	conn.Close()

	consumer, err := queue.NewTransactionConsumer(conn, new(MockRepo))
	assert.NoError(t, err)

	err = consumer.StartConsuming()
	assert.Error(t, err)
}
