package domainorder

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

//CreateOrderCommand is a dto to pass the customer info and the event type, in order to create the command
type CreateOrderCommand struct {
	Customer  string
	EventType eventsourcing.EventType
}

//UpdateOrderCommand is a dto to pass the customer and order info and the event type, in order to create the command
type UpdateOrderCommand struct {
	IDOrder   string
	Customer  string
	EventType eventsourcing.EventType
}

//AddOrderLineCommand is a dto to pass the order info and the event type, in order to create the command
type AddOrderLineCommand struct {
	Price     uint
	Meal      string
	IDOrder   string
	Quantity  uint
	EventType eventsourcing.EventType
}

//UpdateQuantityCommand is a dto to pass the order info and the event type, in order to create the command
type UpdateQuantityCommand struct {
	IDOrderLine string
	Quantity    uint
	EventType   eventsourcing.EventType
}

//DeleteOrderLine is a dto to pass the order info and the event type, in order to create the command
type DeleteOrderLine struct {
	IDOrderLine string
	EventType   eventsourcing.EventType
}

//OrderCommandHandler is a struct use for OrderCommand methods
type OrderCommandHandler struct{}

//Handle handles the order command and pushes the right event to RMQ
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

//NewOrderCommandHandler returns a new OrderCommandHandler
func NewOrderCommandHandler() *OrderCommandHandler {
	return &OrderCommandHandler{}
}

//OrderLineCommandHandler is a struct use for OrderCommandLine methods
type OrderLineCommandHandler struct{}

//Handle handles the order line command and pushes the right event to RMQ
func (ch OrderLineCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
	case *AddOrderLineCommand:
		orderLine := &models.OrderLine{
			Meal:     cmd.Meal,
			Price:    cmd.Price,
			OrderID:  cmd.IDOrder,
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

//NewOrderLineCommandHandler returns a new OrderLineCommandHandler
func NewOrderLineCommandHandler() *OrderLineCommandHandler {
	return &OrderLineCommandHandler{}
}
