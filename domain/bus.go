package domain

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/database"
	domainOrder "github.com/HETIC-MT-P2021/CQRSES_GROUP5/domain/order"
	"log"
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

	_ = CommandBus.RegisterHandler(domainOrder.NewCreateOrderCommandHandler(), &domainOrder.CreateOrderCommand{})
	_ = CommandBus.RegisterHandler(domainOrder.NewUpdateOrderCommandHandler(), &domainOrder.UpdateOrderCommand{})

	_ = CommandBus.RegisterHandler(domainOrder.NewAddOrderLineCommandHandler(), &domainOrder.AddOrderLineCommand{})
	_ = CommandBus.RegisterHandler(domainOrder.NewUpdateQuantityCommandHandler(), &domainOrder.UpdateQuantityCommand{})
	_ = CommandBus.RegisterHandler(domainOrder.NewDeleteOrderLineCommandHandler(), &domainOrder.DeleteOrderLine{})

	_ = QueryBus.RegisterHandler(domainOrder.NewOrderQueryHandler(), &domainOrder.GetOrderQuery{})

	log.Printf("database.EsConn buses: %v", database.EsConn)
}
