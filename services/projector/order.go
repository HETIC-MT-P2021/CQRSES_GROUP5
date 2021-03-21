package projector

import (
	"context"
	"fmt"
	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"github.com/HETIC-MT-P2021/gocqrs/database"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	"github.com/HETIC-MT-P2021/gocqrs/services/elasticsearch"
)

func ProjectEvent(event eventsourcing.Event) error {
	orderRepository := elasticsearch.NewOrderRepository(database.EsConn)
	ctx := context.Background()

	if event.Type == eventsourcing.AddOrder {
		if order, ok := event.Payload.(*models.Order); ok {
			orderRepository.PersistOrder(ctx, order)
		}

		return fmt.Errorf("could not get order from interface")
	}

	if event.Type == eventsourcing.AddOrderLine {
		if order, ok := event.Payload.(*models.OrderLine); ok {
			orderRepository.PersistOrderLine(ctx, order)
		}

		return fmt.Errorf("could not get orderLine from interface")
	}

	if event.Type == eventsourcing.UpdateOrderLine {
		if order, ok := event.Payload.(*models.OrderLine); ok {
			orderRepository.UpdateOrderLine(ctx, order)
		}

		return fmt.Errorf("could not update order from interface")
	}

	return nil
}
