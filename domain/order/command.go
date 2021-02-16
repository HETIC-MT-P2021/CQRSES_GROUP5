package domain_order

import (
	"errors"
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"log"
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
	log.Printf("errrrr")
	switch command.Payload().(type) {
	case *CreateOrderCommand:
		//order := &models.Order{
		//	TotalPrice: 0,
		//	Customer:   cmd.Customer,
		//	Reference:  helpers.RandomString10(),
		//	Lines:      []*models.OrderLine{},
		//}
		//if err := order_repository.PersistOrder(order); err != nil {
		//	return err
		//}

	case *AddOrderLineCommand:
	default:
		log.Printf("YOU SHALL NOT BE ZAIRE")
		return errors.New("bad command type")
	}
	
	
	log.Printf("poukoa tu passes pas laaaaaaaaaa")

	return nil
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}
