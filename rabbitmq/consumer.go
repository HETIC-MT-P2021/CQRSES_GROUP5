package rabbitmq

import (
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/streadway/amqp"
	"log"
)

//StartRBMQConsumer add and register a consumer to the "events" RBMQ Queue
func StartRBMQConsumer() error {
	cfg := Env{}

	if err := env.Parse(&cfg); err != nil {
		return fmt.Errorf("failed to parse env: %v", err)
	}

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s",
		cfg.RabbitMqPass,
		cfg.RabbitMqUser,
		cfg.RabbitMqHost,
		cfg.RabbitMqPort,
	))

	if err != nil {
		return fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		return fmt.Errorf("failed to open a channel: %v", err)
	}

	q, err := ch.QueueDeclare(
		"events", // name
		false,    // durable
		false,    // delete when unused
		false,    // exclusive
		false,    // no-wait
		nil,      // arguments
	)

	if err != nil {
		return fmt.Errorf("failed to declare queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return fmt.Errorf("failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			// Pass event messages to projectors here depending on event type
			log.Printf("received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for event messages. To exit press CTRL+C")
	<-forever

	return nil
}
