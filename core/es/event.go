package es

import (
	"time"
)

// Event is the structure of an event
type Event struct {
	ID        string
	Typology  Typology
	Payload   interface{}
	CreatedAt time.Time
	Index     uint
}

// Typology of an event
type Typology string

// Typology types
const (
	Create Typology = "create"
	Put    Typology = "put"
	Delete Typology = "delete"
)