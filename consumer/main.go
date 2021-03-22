package main

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"github.com/HETIC-MT-P2021/gocqrs/helpers"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	"github.com/HETIC-MT-P2021/gocqrs/services/projector"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/")
	helpers.FailOnError(err, "Failed to connect to RabbitMQ")
	if err != nil {
		log.Fatalf("could not connect to rabbitmq: %v", err)
	}

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("failed to open a channel: %v", err)
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
		log.Fatalf("failed to declare queue: %v", err)
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
		log.Fatalf("failed to register a consumer: %v", err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			//var order models.Order
			buf := bytes.NewBuffer(d.Body)
			dec := gob.NewDecoder(buf)
			e := eventsourcing.Event{}
			gob.Register(models.Order{})
			gob.Register(models.OrderLine{})

			if err := dec.Decode(&e); err != nil {
				log.Fatal(err)
			}

			projector.ProjectEvent(e)
		}
	}()

	log.Printf(" [*] Waiting for event messages. To exit press CTRL+C")
	<-forever
}
