package domainorder

import (
	"errors"
	"fmt"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/eventsourcing"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/services"
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
type CreateOrderCommandHandler struct{}

//Handle handles the order command and pushes the right event to RMQ
func (ch CreateOrderCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
	case *CreateOrderCommand:
		order := models.Order{
			TotalPrice: 0,
			Customer:   cmd.Customer,
			Reference:  helpers.RandomString10(),
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
	default:
		return errors.New("bad command type")
	}

	return nil
}

//NewOrderCommandHandler returns a new OrderCommandHandler
func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}

//OrderCommandHandler is a struct use for OrderCommand methods
type UpdateOrderCommandHandler struct{}

//Handle handles the order command and pushes the right event to RMQ
func (ch UpdateOrderCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
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
func NewUpdateOrderCommandHandler() *UpdateOrderCommandHandler {
	return &UpdateOrderCommandHandler{}
}

//OrderLineCommandHandler is a struct use for OrderCommandLine methods
type AddOrderLineCommandHandler struct{}

//Handle handles the order line command and pushes the right event to RMQ
func (ch AddOrderLineCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
	case *AddOrderLineCommand:
		orderLine := models.OrderLine{
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

	default:
		return errors.New("bad command type")
	}

	return nil
}

//NewOrderLineCommandHandler returns a new OrderLineCommandHandler
func NewAddOrderLineCommandHandler() *AddOrderLineCommandHandler {
	return &AddOrderLineCommandHandler{}
}

//OrderLineCommandHandler is a struct use for OrderCommandLine methods
type UpdateQuantityCommandHandler struct{}

//Handle handles the order line command and pushes the right event to RMQ
func (ch UpdateQuantityCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
	case *UpdateQuantityCommand:
		orderLine := models.OrderLine{
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

	default:
		return errors.New("bad command type")
	}

	return nil
}

//NewOrderLineCommandHandler returns a new OrderLineCommandHandler
func NewUpdateQuantityCommandHandler() *UpdateQuantityCommandHandler {
	return &UpdateQuantityCommandHandler{}
}

//OrderLineCommandHandler is a struct use for OrderCommandLine methods
type DeleteOrderLineCommandHandler struct{}

//Handle handles the order line command and pushes the right event to RMQ
func (ch DeleteOrderLineCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
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
func NewDeleteOrderLineCommandHandler() *DeleteOrderLineCommandHandler {
	return &DeleteOrderLineCommandHandler{}
}

