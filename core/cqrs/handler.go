package cqrs

//CommandHandler is an interface for all command handlers methods
type CommandHandler interface {
	Handle(message CommandMessage) error
}

//QueryHandler is an interface for all query handlers methods
type QueryHandler interface {
	Handle(message QueryMessage) error
}
