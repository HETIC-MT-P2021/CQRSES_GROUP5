package elastic_search

import (
	"context"
	"github.com/HETIC-MT-P2021/gocqrs/models"
)

const (
	OrderIndex     = "index_order"
	OrderEventIndex = "index_order_event"
	OrderLineIndex = "index_order_line"
)

type OrderRepository struct {
	EsConnector *EsConnector
}

//GetOrder is for getting an order
func (repository *OrderRepository) GetOrder(ctx context.Context, orderID string) (*Document, error) {
	return repository.EsConnector.GetDocumentByIndexAndID(ctx, OrderIndex, orderID)
}

//GetOrderLine is for getting an order
func (repository *OrderRepository)  GetOrderLine(ctx context.Context, orderLineID string) (*Document, error) {
	return repository.EsConnector.GetDocumentByIndexAndID(ctx, OrderLineIndex, orderLineID)
}

//PersistOrder persists the order as is in elastic search
func (repository *OrderRepository) PersistOrder(ctx context.Context, order *models.Order) error {

	return repository.EsConnector.NewDocument(ctx, OrderIndex, &Document{
		ID:   order.ID,
		Body: order,
	})
}

//PersistOrderLine persists a fixed state of orderLine
func (repository *OrderRepository) PersistOrderLine(ctx context.Context, orderLine *models.OrderLine) error {
	
	return repository.EsConnector.NewDocument(ctx, OrderLineIndex, &Document{
		ID:   orderLine.ID,
		Body: orderLine,
	})
}

func (repository *OrderRepository) UpdateOrderLine(ctx context.Context, orderLine *models.OrderLine) error{
	
	_, err := repository.EsConnector.UpdateDocument(ctx, OrderLineIndex, &Document{
		ID:   orderLine.ID,
		Body: orderLine,
	})
	
	return err
}
