package cqrs

type CommandHandler interface {
	Handle(message CommandMessage) error
}

type QueryHandler interface {
	Handle(message QueryMessage) error
}
