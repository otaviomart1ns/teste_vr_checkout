package queue

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQPublisher struct {
	channel *amqp.Channel
	queue   string
}

func NewRabbitMQPublisher(conn *amqp.Connection, queueName string) (*RabbitMQPublisher, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &RabbitMQPublisher{
		channel: ch,
		queue:   queueName,
	}, nil
}

func (p *RabbitMQPublisher) PublishTransaction(body []byte) error {
	return p.channel.Publish(
		"",          // exchange
		p.queue,     // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}