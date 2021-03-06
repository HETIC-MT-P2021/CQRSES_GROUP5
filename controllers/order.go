package controllers

import (
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"github.com/HETIC-MT-P2021/gocqrs/domain"
	domain_order "github.com/HETIC-MT-P2021/gocqrs/domain/order"
	"github.com/HETIC-MT-P2021/gocqrs/helpers"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	"net/http"
)

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
