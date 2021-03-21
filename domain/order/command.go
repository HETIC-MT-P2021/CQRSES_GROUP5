package domain_order

import (
	"errors"
	"fmt"
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"github.com/HETIC-MT-P2021/gocqrs/helpers"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	"github.com/HETIC-MT-P2021/gocqrs/services"
	"time"
)

type CreateOrderCommand struct {
	Customer  string
	EventType eventsourcing.EventType
}

type UpdateOrderCommand struct {
	IDOrder   uint
	Customer  string
	EventType eventsourcing.EventType
}

type AddOrderLineCommand struct {
	Price     uint
	Meal      string
	IDOrder   uint
	Quantity  uint
	EventType eventsourcing.EventType
}

type UpdateQuantityCommand struct {
	IDOrderLine uint
	Quantity    uint
	EventType   eventsourcing.EventType
}

type DeleteOrderLine struct {
	IDOrderLine uint
	EventType   eventsourcing.EventType
}

type OrderCommandHandler struct{}

func (ch OrderCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
	case *CreateOrderCommand:
		order := &models.Order{
			TotalPrice: 0,
			Customer:   cmd.Customer,
			Reference:  helpers.RandomString10(),
			Lines:      []*models.OrderLine{},
		}

		// Creates and send an Event to RabbitMQ
		createOrderEvent := eventsourcing.Event{
			Type:           cmd.EventType,
			Payload:        order,
			CreatedAt:      time.Time{},
			AggregateIndex: 1, // Order aggregation Index
		}

		err := services.PublishEventToRBMQ(createOrderEvent)

		if err != nil {
			return fmt.Errorf("failed to publish an event: %v", err)
		}

	case *UpdateOrderCommand:
		order := &models.Order{
			ID:       cmd.IDOrder,
			Customer: cmd.Customer,
		}

		// Creates and send an Event to RabbitMQ
		updateOrderEvent := eventsourcing.Event{
			Type:           cmd.EventType,
			Payload:        order,
			CreatedAt:      time.Time{},
			AggregateIndex: 1, // Order aggregation Index
		}

		err := services.PublishEventToRBMQ(updateOrderEvent)

		if err != nil {
			return fmt.Errorf("failed to publish an event: %v", err)
		}

	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewOrderCommandHandler() *OrderCommandHandler {
	return &OrderCommandHandler{}
}

type OrderLineCommandHandler struct{}

func (ch OrderLineCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
	case *AddOrderLineCommand:
		orderLine := &models.OrderLine{
			Meal:     cmd.Meal,
			Price:    cmd.Price,
			IDOrder:  cmd.IDOrder,
			Quantity: cmd.Quantity,
		}

		// Creates and send an Event to RabbitMQ
		addOrderLineEvent := eventsourcing.Event{
			Type:           cmd.EventType,
			Payload:        orderLine,
			CreatedAt:      time.Time{},
			AggregateIndex: 1, // Order aggregation Index
		}

		err := services.PublishEventToRBMQ(addOrderLineEvent)

		if err != nil {
			return fmt.Errorf("failed to publish an event: %v", err)
		}

	case *UpdateQuantityCommand:
		orderLine := &models.OrderLine{
			ID:       cmd.IDOrderLine,
			Quantity: cmd.Quantity,
		}

		// Creates and send an Event to RabbitMQ
		updateQuantityEvent := eventsourcing.Event{
			Type:           cmd.EventType,
			Payload:        orderLine,
			CreatedAt:      time.Time{},
			AggregateIndex: 1, // Order aggregation Index
		}

		err := services.PublishEventToRBMQ(updateQuantityEvent)

		if err != nil {
			return fmt.Errorf("failed to publish an event: %v", err)
		}

	case *DeleteOrderLine:
		// Creates and send an Event to RabbitMQ
		updateQuantityEvent := eventsourcing.Event{
			Type:           cmd.EventType,
			Payload:        cmd.IDOrderLine,
			CreatedAt:      time.Time{},
			AggregateIndex: 1, // Order aggregation Index
		}

		err := services.PublishEventToRBMQ(updateQuantityEvent)

		if err != nil {
			return fmt.Errorf("failed to publish an event: %v", err)
		}

	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewOrderLineCommandHandler() *OrderLineCommandHandler {
	return &OrderLineCommandHandler{}
}
