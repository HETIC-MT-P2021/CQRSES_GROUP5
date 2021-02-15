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

// Type of an event
type EventType string

// EventType types
const (
	AddOrder        EventType = "addOrder"
	UpdateQuantity  EventType = "updateQuantity"
	AddOrderLine    EventType = "addOrderLine"
	UpdateOrderLine EventType = "updateOrderLine"
)
