package controllers

import (
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"github.com/HETIC-MT-P2021/gocqrs/domain"
	domain_order "github.com/HETIC-MT-P2021/gocqrs/domain/order"
	"github.com/HETIC-MT-P2021/gocqrs/helpers"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

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
	orderID, err := helpers.ParseUInt(muxVars["id"])

	if err != nil {
		log.Printf("could not parse id into int: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not parse id")
		return
	}

	command := cqrs.NewCommandMessage(&domain_order.UpdateOrderCommand{
		IDOrder:   orderID,
		Customer:  order.Customer,
		EventType: eventsourcing.UpdateOrder,
	})

	err = domain.CommandBus.Dispatch(command)

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
	orderID, err := helpers.ParseUInt(muxVars["id"])

	if err != nil {
		log.Printf("could not parse id into int: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not parse id")
		return
	}

	command := cqrs.NewCommandMessage(&domain_order.AddOrderLineCommand{
		IDOrder:   orderID,
		Price:     orderLine.Price,
		Meal:      orderLine.Meal,
		EventType: eventsourcing.AddOrderLine,
	})

	err = domain.CommandBus.Dispatch(command)

	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, command)
}

//UpdateOrderLineQuantity creates a new UpdateOrderLineQuantity command
func UpdateOrderLineQuantity(w http.ResponseWriter, r *http.Request) {
	muxVars := mux.Vars(r)
	orderLineID, err := helpers.ParseUInt(muxVars["id"])

	if err != nil {
		log.Printf("could not parse id into int: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not parse id")
		return
	}

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
	orderLineID, err := helpers.ParseUInt(muxVars["id"])

	if err != nil {
		log.Printf("could not parse id into int: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "could not parse id")
		return
	}

	command := cqrs.NewCommandMessage(&domain_order.DeleteOrderLine{
		IDOrderLine: orderLineID,
		EventType:   eventsourcing.DeleteOrderLine,
	})

	err = domain.CommandBus.Dispatch(command)

	if err != nil {
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.WriteJSON(w, http.StatusOK, command)
}
