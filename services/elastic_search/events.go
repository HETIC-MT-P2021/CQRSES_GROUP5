package elastic_search

import (
	"context"
	"fmt"
	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
)

func (repository *OrderRepository) AddEventInES(ctx context.Context, event *eventsourcing.Event) error {
	
	return repository.EsConnector.NewDocument(ctx, OrderEventIndex, &Document{
		ID:   fmt.Sprintf("%s-%s", event.Type, event.ID),
		Body: event,
	})
}
