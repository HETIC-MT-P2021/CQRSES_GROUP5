package domainorder

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/database"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/services/elasticsearch"
)

//GetOrderQuery is a dto to pass the order id, in order to create the query for order
type GetOrderQuery struct {
	OrderID string
}

//OrderQueryHandler is a struct use for OrderQuery methods
type OrderQueryHandler struct{}

//Handle handles the order query and retrieves the order from Elastic Search
func (ch OrderQueryHandler) Handle(query cqrs.QueryMessage, w *http.ResponseWriter) error {

	ctx := context.Background()

	switch quy := query.Payload().(type) {
	case *GetOrderQuery:
		orderRepository := elasticsearch.NewOrderRepository(database.EsConn)

		orderDoc, err := orderRepository.GetOrder(ctx, quy.OrderID)
		if err != nil {
			return fmt.Errorf("failed to retrieve order: %v", err)
		}

		helpers.WriteJSON(*w, http.StatusOK, orderDoc)
		return err

	default:
		return errors.New("bad query type")
	}
}

//NewOrderQueryHandler returns a new OrderQueryHandler
func NewOrderQueryHandler() *OrderQueryHandler {
	return &OrderQueryHandler{}
}
