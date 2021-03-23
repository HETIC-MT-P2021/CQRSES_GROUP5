package domain

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/cqrs"
	domain_order "github.com/HETIC-MT-P2021/CQRSES_GROUP5/domain/order"
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

	_ = CommandBus.RegisterHandler(domain_order.NewCreateOrderCommandHandler(), &domain_order.CreateOrderCommand{})
	_ = CommandBus.RegisterHandler(domain_order.NewUpdateOrderCommandHandler(), &domain_order.UpdateOrderCommand{})

	_ = CommandBus.RegisterHandler(domain_order.NewAddOrderLineCommandHandler(), &domain_order.AddOrderLineCommand{})
	_ = CommandBus.RegisterHandler(domain_order.NewUpdateQuantityCommandHandler(), &domain_order.UpdateQuantityCommand{})
	_ = CommandBus.RegisterHandler(domain_order.NewDeleteOrderLineCommandHandler(), &domain_order.DeleteOrderLine{})

	_ = QueryBus.RegisterHandler(domain_order.NewOrderQueryHandler(), &domain_order.GetOrderQuery{})
}
