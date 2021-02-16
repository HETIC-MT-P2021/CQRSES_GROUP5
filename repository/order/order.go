package order_repository

import (
	"context"
	"fmt"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	"github.com/HETIC-MT-P2021/gocqrs/services"
)

const (
	OrderIndex     = "index_order"
	OrderLineIndex = "index_order_line"
)

//GetOrder is for getting an order
func GetOrder(ctx context.Context, orderID string) (*models.Order, error) {
	document, err := services.GetDocumentByIndexAndID(ctx, OrderIndex, orderID)
	if err != nil {
		return nil, fmt.Errorf("could not get order : %v", err)
	}

	order, ok := document.Body.(*models.Order)
	if !ok {
		return nil, fmt.Errorf("could not format data: %v", ok)
	}

	return order, nil
}

//GetOrderLine is for getting an order
func GetOrderLine(ctx context.Context, orderLineID string) (*models.OrderLine, error) {
	document, err := services.GetDocumentByIndexAndID(ctx, OrderLineIndex, orderLineID)
	if err != nil {
		return nil, fmt.Errorf("could not get order line: %v", err)
	}

	orderLine, ok := document.Body.(*models.OrderLine)
	if !ok {
		return nil, fmt.Errorf("could not format data: %v", ok)
	}

	return orderLine, nil
}

//CreateOrderIndex creates the Order index in Elastic Search
func CreateOrderIndex(ctx context.Context) error {

	return services.NewIndex(ctx, OrderIndex)
}

//CreateOrderLineIndex creates the Order Line index in Elastic Search
func CreateOrderLineIndex(ctx context.Context) error {
	return services.NewIndex(ctx, OrderLineIndex)
}

// TODO : Create an event and store it to ElasticSearch
func PersistOrder(ctx context.Context, order *models.Order) error {
	document := &services.Document{
		ID : order.ID,
		Body: order,
	}

	return services.NewDocument(ctx, OrderIndex, document)
}

// TODO : Create an event and store it to ElasticSearch
func PersistOrderLine(ctx context.Context, orderLine *models.OrderLine) error {
	document := &services.Document{
		Body: orderLine,
	}

	return services.NewDocument(ctx, OrderIndex, document)
}
