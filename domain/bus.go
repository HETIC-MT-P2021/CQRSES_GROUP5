package domain

import (
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	domain_order "github.com/HETIC-MT-P2021/gocqrs/domain/order"
)

var CommandBus *cqrs.CommandBus
var QueryBus *cqrs.QueryBus

func InitBusses() {
	CommandBus = cqrs.NewCommandBus()
	QueryBus = cqrs.NewQueryBus()

	_ = CommandBus.RegisterHandler(domain_order.NewCreateOrderCommandHandler(), &domain_order.CreateOrderCommand{})
}
