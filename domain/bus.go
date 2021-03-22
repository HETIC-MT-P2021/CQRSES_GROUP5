package domain

import (
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	domain_order "github.com/HETIC-MT-P2021/gocqrs/domain/order"
)

// Declaration of CQRS buses
var (
	CommandBus *cqrs.CommandBus
	QueryBus   *cqrs.QueryBus
)

//InitBusses inits all command and queries handlers
func InitBusses() {
	CommandBus = cqrs.NewCommandBus()
	QueryBus = cqrs.NewQueryBus()

	_ = CommandBus.RegisterHandler(domain_order.NewOrderCommandHandler(), &domain_order.CreateOrderCommand{})
	_ = CommandBus.RegisterHandler(domain_order.NewOrderCommandHandler(), &domain_order.UpdateOrderCommand{})

	_ = CommandBus.RegisterHandler(domain_order.NewOrderLineCommandHandler(), &domain_order.AddOrderLineCommand{})
	_ = CommandBus.RegisterHandler(domain_order.NewOrderLineCommandHandler(), &domain_order.UpdateQuantityCommand{})
	_ = CommandBus.RegisterHandler(domain_order.NewOrderLineCommandHandler(), &domain_order.DeleteOrderLine{})

	_ = QueryBus.RegisterHandler(domain_order.NewOrderQueryHandler(), &domain_order.GetOrderQuery{})
}
