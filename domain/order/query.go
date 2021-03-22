package domainorder

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	"github.com/HETIC-MT-P2021/gocqrs/database"
	"github.com/HETIC-MT-P2021/gocqrs/helpers"
	"github.com/HETIC-MT-P2021/gocqrs/services/elasticsearch"
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

		log.Printf("id : %s", quy.OrderID)

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
