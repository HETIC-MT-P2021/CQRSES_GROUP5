package controllers

import (
<<<<<<< Updated upstream
	"github.com/HETIC-MT-P2021/gocqrs/core/cqrs"
	"github.com/HETIC-MT-P2021/gocqrs/core/eventsourcing"
	"github.com/HETIC-MT-P2021/gocqrs/domain"
	domain_order "github.com/HETIC-MT-P2021/gocqrs/domain/order"
=======
>>>>>>> Stashed changes
	"github.com/HETIC-MT-P2021/gocqrs/helpers"
	"github.com/HETIC-MT-P2021/gocqrs/models"
	order_repository "github.com/HETIC-MT-P2021/gocqrs/repository/order"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	
	order := models.Order{}
	
	if err := helpers.ReadJSON(w, r, &order); err != nil {
		log.Printf("could not read request: %v", err)
		return
	}
	
	ctx := r.Context()
	
	//err := order_repository.CreateOrderIndex(ctx)
	//if err != nil {
	//	log.Printf("could not create order index: %v", err)
	//	helpers.WriteErrorJSON(w, http.StatusInternalServerError, "error")
	//	return 
	//}
	
	if err := order_repository.PersistOrder(ctx, &order); err != nil {
		log.Printf("could not create order: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "error")
		return
	}

<<<<<<< Updated upstream
	command := cqrs.NewCommandMessage(&domain_order.CreateOrderCommand{
		Customer:  order.Customer,
		EventType: eventsourcing.AddOrder,
	})
=======
	//command := cqrs.NewCommandMessage(&domain_order.CreateOrderCommand{Order: &order})
	//
	//if err := domain.CommandBus.Dispatch(command); err != nil {
	//	helpers.WriteErrorJSON(w, http.StatusInternalServerError, err.Error())
	//	return
	//}

	helpers.WriteJSON(w, http.StatusOK, order)
}


func GetOrder(w http.ResponseWriter, r *http.Request) {
	
	muxVars := mux.Vars(r)
	id := muxVars["order_id"]
>>>>>>> Stashed changes

	ctx := r.Context()

	order, err := order_repository.GetOrder(ctx, id)
	if err != nil {
		log.Printf("error fetching order: %v", err)
		helpers.WriteErrorJSON(w, http.StatusInternalServerError, "error")
		return
	}

	helpers.WriteJSON(w, http.StatusOK, order)
}
