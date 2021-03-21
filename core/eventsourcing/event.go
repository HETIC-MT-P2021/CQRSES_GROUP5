package eventsourcing

import (
	"time"
)

// Event is the structure of an event
type Event struct {
	ID             string
	Type           EventType
	Payload        interface{}
	CreatedAt      time.Time
	AggregateIndex uint
}

// EventType is the type of event
type EventType string

// EventType types
const (
	AddOrder        EventType = "addOrder"
	UpdateOrder     EventType = "updateOrder"
	UpdateQuantity  EventType = "updateQuantity"
	AddOrderLine    EventType = "addOrderLine"
	DeleteOrderLine EventType = "deleteOrderLine"
	UpdateOrderLine EventType = "updateOrderLine"
)
