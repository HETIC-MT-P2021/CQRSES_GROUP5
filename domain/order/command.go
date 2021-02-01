package domain_order

import (
	"errors"
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	"github.com/HETIC-MT-P2021/gocqrs/database"
	"github.com/HETIC-MT-P2021/gocqrs/helpers"
	"github.com/HETIC-MT-P2021/gocqrs/models"
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
	db := database.DbConn
	repository := models.Repository{Conn: db}

	switch cmd := command.Payload().(type) {
	case *CreateOrderCommand:
		order := &models.Order{
			TotalPrice: 230,
			Client:     cmd.Client,
			Reference:  helpers.RandomString10(),
			Lines:      []*models.OrderLine{},
		}
		order.NewOrder(order)
	case *AddOrderLineCommand:
		orderLine := &models.OrderLine{
			Meal:    cmd.Meal,
			Price:   cmd.Price,
			IDOrder: cmd.IDOrder,
		}
		repository.AddOrderLine(orderLine)
	default:
		return errors.New("bad command type")
	}

	return nil
}

func NewCreateOrderCommandHandler() *CreateOrderCommandHandler {
	return &CreateOrderCommandHandler{}
}
