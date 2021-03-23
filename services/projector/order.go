package projector

import (
	"context"
	"fmt"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/eventsourcing"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/models"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/services/elasticsearch"
	"time"
)

//ProjectEvent project an event into a readModel depending on the event type
func ProjectEvent(ctx context.Context, event eventsourcing.Event) error {
	
	esConn, err := database.GetEsConn(ctx, 5 * time.Second)
	if err != nil {
		return fmt.Errorf("could not get es connection : %v", err)
	}
	
	orderRepository := elasticsearch.NewOrderRepository(esConn)

	switch event.Type {

	case eventsourcing.AddOrder:
		if order, ok := event.Payload.(models.Order); ok {
			if err := orderRepository.PersistOrder(ctx, &order); err != nil {
				return fmt.Errorf("err : %v", err)
			}
		}

		return fmt.Errorf("could not get order from interface")

	case eventsourcing.AddOrderLine:
		if order, ok := event.Payload.(models.OrderLine); ok {
			if err := orderRepository.PersistOrderLine(ctx, &order); err != nil {
				return fmt.Errorf("err : %v", err)
			}
		}

		return fmt.Errorf("could not get orderLine from interface")

	case eventsourcing.UpdateOrderLine:
		if order, ok := event.Payload.(models.OrderLine); ok {
			if err := orderRepository.UpdateOrderLine(ctx, &order); err != nil {
				return fmt.Errorf("err : %v", err)
			}
		}

		return fmt.Errorf("could not update order from interface")
	}

	return nil
}
