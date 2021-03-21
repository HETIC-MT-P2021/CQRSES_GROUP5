package projector

import "github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"

func ProjectEvent(event eventsourcing.Event) error {

	if event.Type == eventsourcing.AddOrder {

	}

	if event.Type == eventsourcing.AddOrderLine {

	}

	if event.Type == eventsourcing.UpdateOrderLine {

	}

	if event.Type == eventsourcing.UpdateQuantity {

	}

	return nil
}
