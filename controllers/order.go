package controllers

import (
	"log"
	"net/http"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/cqrs"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/core/eventsourcing"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/domain"
	domain_order "github.com/HETIC-MT-P2021/CQRSES_GROUP5/domain/order"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/helpers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP5/models"
	"github.com/gorilla/mux"
)

//GetOrder gets an order from ES database
func GetOrder(w http.ResponseWriter, r *http.Request) {

	muxVars := mux.Vars(r)
	orderID := muxVars["id"]

	query := cqrs.NewQueryMessage(&domain_order.GetOrderQuery{OrderID: orderID})

	err := domain.QueryBus.Dispatch(query, &w)
	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, "hellow")
}

//CreateOrder creates a new CreateOrder command (CQRS pattern)
func CreateOrder(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}

	if err := helpers.ReadJSON(w, r, &order); err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "can not parse JSON body")
		return
	}

	command := cqrs.NewCommandMessage(&domain_order.CreateOrderCommand{
		Customer:  order.Customer,
		EventType: eventsourcing.AddOrder,
	})

	err := domain.CommandBus.Dispatch(command)

	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, command)
}

//UpdateOrder creates a new UpdateOrder command (CQRS pattern)
func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	order := models.Order{}

	if err := helpers.ReadJSON(w, r, &order); err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "can not parse JSON body")
		return
	}

	muxVars := mux.Vars(r)
	orderID := muxVars["id"]

	command := cqrs.NewCommandMessage(&domain_order.UpdateOrderCommand{
		IDOrder:   orderID,
		Customer:  order.Customer,
		EventType: eventsourcing.UpdateOrder,
	})

	err := domain.CommandBus.Dispatch(command)
	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, command)
}

//AddOrderLine creates a new AddOrderLine command (CQRS pattern)
func AddOrderLine(w http.ResponseWriter, r *http.Request) {
	orderLine := models.OrderLine{}

	if err := helpers.ReadJSON(w, r, &orderLine); err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "can not parse JSON body")
		return
	}

	muxVars := mux.Vars(r)
	orderID := muxVars["id"]

	command := cqrs.NewCommandMessage(&domain_order.AddOrderLineCommand{
		IDOrder:   orderID,
		Price:     orderLine.Price,
		Meal:      orderLine.Meal,
		EventType: eventsourcing.AddOrderLine,
	})

	err := domain.CommandBus.Dispatch(command)
	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, command)
}

//UpdateOrderLineQuantity creates a new UpdateOrderLineQuantity command
func UpdateOrderLineQuantity(w http.ResponseWriter, r *http.Request) {
	muxVars := mux.Vars(r)
	orderLineID := muxVars["id"]

	orderLineQuantity, err := helpers.ParseUInt(muxVars["quantity"])

	if err != nil {
		log.Printf("could not parse quantity into int: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not parse id")
		return
	}

	command := cqrs.NewCommandMessage(&domain_order.UpdateQuantityCommand{
		IDOrderLine: orderLineID,
		Quantity:    orderLineQuantity,
		EventType:   eventsourcing.UpdateQuantity,
	})

	err = domain.CommandBus.Dispatch(command)

	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, command)
}

//DeleteOrderLine creates a new DeleteOrderLine command
func DeleteOrderLine(w http.ResponseWriter, r *http.Request) {
	muxVars := mux.Vars(r)
	orderLineID := muxVars["id"]

	command := cqrs.NewCommandMessage(&domain_order.DeleteOrderLine{
		IDOrderLine: orderLineID,
		EventType:   eventsourcing.DeleteOrderLine,
	})

	err := domain.CommandBus.Dispatch(command)

	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, command)
}
