package elasticsearch

import (
	"context"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/models"
	"github.com/olivere/elastic/v7"
	"github.com/google/uuid"
)

// Indexes for ES
const (
	OrderIndex      = "index_order"
	OrderEventIndex = "index_order_event"
	OrderLineIndex  = "index_order_line"
)

//OrderRepository encapsulates the ESConnector for all repository methods
type OrderRepository struct {
	EsConnector *EsConnector
}

//NewOrderRepository return a new OrderRepository
func NewOrderRepository(client *elastic.Client) *OrderRepository {
	return &OrderRepository{EsConnector: NewEsConnector(client)}
}

//GetOrder is for getting an order
func (repository *OrderRepository) GetOrder(ctx context.Context, orderID string) (*Document, error) {
	return repository.EsConnector.GetDocumentByIndexAndID(ctx, OrderIndex, orderID)
}

//GetOrderLine is for getting an order
func (repository *OrderRepository) GetOrderLine(ctx context.Context, orderLineID string) (*Document, error) {
	return repository.EsConnector.GetDocumentByIndexAndID(ctx, OrderLineIndex, orderLineID)
}

//PersistOrder persists the order as is in elastic search
func (repository *OrderRepository) PersistOrder(ctx context.Context, order *models.Order) error {
	order.ID = uuid.New().String()
	
	return repository.EsConnector.NewDocument(ctx, OrderIndex, &Document{
		ID:   order.ID,
		Body: order,
	})
}

//PersistOrderLine persists a fixed state of orderLine
func (repository *OrderRepository) PersistOrderLine(ctx context.Context, orderLine *models.OrderLine) error {
	orderLine.ID = uuid.New().String()
	
	return repository.EsConnector.NewDocument(ctx, OrderLineIndex, &Document{
		ID:   orderLine.ID,
		Body: orderLine,
	})
}

//UpdateOrderLine updates an order line in ES
func (repository *OrderRepository) UpdateOrderLine(ctx context.Context, orderLine *models.OrderLine) error {

	_, err := repository.EsConnector.UpdateDocument(ctx, OrderLineIndex, &Document{
		ID:   orderLine.ID,
		Body: orderLine,
	})

	return err
}
