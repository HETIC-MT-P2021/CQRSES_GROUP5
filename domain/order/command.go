package domain_order

import (
	"errors"
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	"github.com/HETIC-MT-P2021/gocqrs/helpers"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	order_repository "github.com/HETIC-MT-P2021/gocqrs/repository/order"
)

type CreateOrderCommand struct {
	Client string
}

type AddOrderLineCommand struct {
	Price   uint
	Meal    string
	IDOrder uint
}

type CreateOrderCommandHandler struct{}

func (ch CreateOrderCommandHandler) Handle(command cqrs.CommandMessage) error {
	switch cmd := command.Payload().(type) {
	case *CreateOrderCommand:
		order := &models.Order{
			TotalPrice: 0,
			Client:     cmd.Client,
			Reference:  helpers.RandomString10(),
			Lines:      []*models.OrderLine{},
		}
		if err := order_repository.PersistOrder(order); err != nil {
			return err
		}

	case *AddOrderLineCommand:
		orderLine := &models.OrderLine{
			Meal:    cmd.Meal,
			Price:   cmd.Price,
			IDOrder: cmd.IDOrder,
		}
		order_repository.PersistOrderLine(orderLine)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}
