package main

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/eventsourcing"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/services/projector"
	"github.com/streadway/amqp"
	"log"
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

	//cfg := &config.Config{
	//	URL:         "http://127.0.0.1:9200",
	//}
	//
	//client, err := elastic.DialWithConfig(context.Background(), cfg)
	//if err != nil {
	//	log.Printf("err : %v", err)
	//	return
	//}

	//log.Printf("client: %+v", client)

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

	log.Printf("database.EsConn from cons: %v", database.EsConn)

	ctx := context.Background()

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

			log.Printf("event : %+v", e)

			if err := projector.ProjectEvent(ctx, e); err != nil {
				log.Fatalf("err on projecting event : %v", err)
			}
		}
	}()

	log.Printf(" [*] Waiting for event messages. To exit press CTRL+C")
	<-forever
}
