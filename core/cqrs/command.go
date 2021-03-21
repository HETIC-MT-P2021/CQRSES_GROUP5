package cqrs

import "fmt"

//CommandMessage is a message command implementation interface
type CommandMessage interface {
	Payload() interface{}
	CommandType() string
}

//CommandBus is a command bus (CQRS pattern)
type CommandBus struct {
	handlers map[string]CommandHandler
}

//NewCommandBus creates a new command bus (CQRS pattern)
func NewCommandBus() *CommandBus {
	cBus := &CommandBus{
		handlers: make(map[string]CommandHandler),
	}

	return cBus
}

//Dispatch dispatches command buses (CQRS pattern)
func (b *CommandBus) Dispatch(command CommandMessage) error {
	if handler, ok := b.handlers[command.CommandType()]; ok {
		return handler.Handle(command)
	}

	return fmt.Errorf("the command bus does not have a handler for command of type: %s", command)
}

//RegisterHandler registers handlers for commands (CQRS pattern)
func (b *CommandBus) RegisterHandler(handler CommandHandler, command interface{}) error {
	typeName := typeOf(command)
	if _, ok := b.handlers[typeName]; ok {
		return fmt.Errorf("duplicate command handler registration with command bus for command of type: %s", typeName)
	}

	b.handlers[typeName] = handler

	return nil
}

//CommandDescriptor encapsulates commands, which can be any type (CQRS pattern)
type CommandDescriptor struct {
	command interface{}
}

//NewCommandMessage creates a new command message (CQRS pattern)
func NewCommandMessage(command interface{}) *CommandDescriptor {
	return &CommandDescriptor{
		command: command,
	}
}

//CommandType returns the type of command
func (c *CommandDescriptor) CommandType() string {
	return typeOf(c.command)
}

//Payload returns the actual command payload of the message.
func (c *CommandDescriptor) Payload() interface{} {
	return c.command
}
