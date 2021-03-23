package cqrs

import (
	"fmt"
	"net/http"
)

//QueryMessage is an interface for all query messages methodss
type QueryMessage interface {
	Payload() interface{}
	QueryType() string
}

//QueryBus encapsulates multiple query handlers
type QueryBus struct {
	handlers map[string]QueryHandler
}

//NewQueryBus returns a new QueryBus
func NewQueryBus() *QueryBus {
	cBus := &QueryBus{
		handlers: make(map[string]QueryHandler),
	}

	return cBus
}

//Dispatch dispatches query buses (CQRS pattern)
func (b *QueryBus) Dispatch(query QueryMessage, w *http.ResponseWriter) error {
	if handler, ok := b.handlers[query.QueryType()]; ok {
		return handler.Handle(query, w)
	}
	return fmt.Errorf("the query bus does not have a handler for query of type: %s", query)
}

//RegisterHandler registers handlers for queries (CQRS pattern)
func (b *QueryBus) RegisterHandler(handler QueryHandler, query interface{}) error {
	typeName := typeOf(query)
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("duplicate query handler registration with query bus for query of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

//QueryDescriptor encapsulates queries, which can be any type (CQRS pattern)
type QueryDescriptor struct {
	query interface{}
}

//NewQueryMessage creates a new query message (CQRS pattern)
func NewQueryMessage(query interface{}) *QueryDescriptor {
	return &QueryDescriptor{
		query: query,
	}
}

//QueryType returns the type of query
func (c *QueryDescriptor) QueryType() string {
	return typeOf(c.query)
}

//Payload returns the actual query payload of the message.
func (c *QueryDescriptor) Payload() interface{} {
	return c.query
}
