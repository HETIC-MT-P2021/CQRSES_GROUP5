package order_repository

import (
	"github.com/HETIC-MT-P2021/gocqrs/models"
)

// TODO
//GetOrder is for getting an order
func GetOrder() (*models.Order, error) {
	return nil, nil
}

// TODO : Create an event and store it to ElasticSearch
func PersistOrder(order *models.Order) error {
	return nil
}

// TODO : Create an event and store it to ElasticSearch
func PersistOrderLine(order *models.OrderLine) error {
	return nil
}

// TODO
func AddOrderLine(line *models.OrderLine) error {
	return nil
}
