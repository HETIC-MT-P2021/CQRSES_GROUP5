package services

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"github.com/HETIC-MT-P2021/gocqrs/rabbitmq"
	"github.com/streadway/amqp"
)

//PublishEventToRBMQ publishes an event sourcing event to RBMQ
func PublishEventToRBMQ(event eventsourcing.Event) error {

	rbmqChanCreation := rabbitmq.RBMQQueuecreation{
		RabbitMQChan:  rabbitmq.RabbitMQChan,
		RabbitMQQueue: rabbitmq.RabbitMQQueue,
	}

	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(event)
	if err != nil {
		log.Fatal("event encode error:", err)
	}

	err = rbmqChanCreation.RabbitMQChan.Publish(
		"",                                  // exchange
		rbmqChanCreation.RabbitMQQueue.Name, // routing key
		false,                               // mandatory
		false,                               // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        network.Bytes(),
		})

	return err
}
