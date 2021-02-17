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

type AddOrderLineCommand struct {
	Price     uint
	Meal      string
	IDOrder   uint
	EventType eventsourcing.EventType
}

type CreateOrderCommandHandler struct{}

func (ch CreateOrderCommandHandler) Handle(command cqrs.CommandMessage) error {
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

	case *AddOrderLineCommand:
		//orderLine := &models.OrderLine{
		//	Meal:    cmd.Meal,
		//	Price:   cmd.Price,
		//	IDOrder: cmd.IDOrder,
		//}
		//order_repository.PersistOrderLine(orderLine)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}
