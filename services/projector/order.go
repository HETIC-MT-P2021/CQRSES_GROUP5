package projector

import (
	"context"
	"fmt"

	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"github.com/HETIC-MT-P2021/gocqrs/database"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	"github.com/HETIC-MT-P2021/gocqrs/services/elasticsearch"
)

//ProjectEvent project an event into a readModel depending on the event type
func ProjectEvent(event eventsourcing.Event) error {
	orderRepository := elasticsearch.NewOrderRepository(database.EsConn)
	ctx := context.Background()
	if event.Type == eventsourcing.AddOrder {
		if order, ok := event.Payload.(models.Order); ok {
			if err := orderRepository.PersistOrder(ctx, &order); err != nil {
				return fmt.Errorf("err : %v", err)
			}
		}

		return fmt.Errorf("could not get order from interface")
	}

	if event.Type == eventsourcing.AddOrderLine {
		if order, ok := event.Payload.(models.OrderLine); ok {
			if err := orderRepository.PersistOrderLine(ctx, &order); err != nil {
				return fmt.Errorf("err : %v", err)
			}
		}

		return fmt.Errorf("could not get orderLine from interface")
	}

	if event.Type == eventsourcing.UpdateOrderLine {
		if order, ok := event.Payload.(models.OrderLine); ok {
			if err := orderRepository.UpdateOrderLine(ctx, &order); err != nil {
				return fmt.Errorf("err : %v", err)
			}
		}

		return fmt.Errorf("could not update order from interface")
	}

	return nil
}
