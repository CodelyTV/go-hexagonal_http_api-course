package command

import "context"

// Bus defines the expected behaviour from a command bus.
type Bus interface {
	// Dispatch is the method used to dispatch new commands.
	Dispatch(context.Context, Command) error
	// Register is the method used to register a new command handler.
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

// Type represents an application command type.
type Type string

// Command represents an application command.
type Command interface {
	Type() Type
}

// Handler defines the expected behaviour from a command handler.
type Handler interface {
	Handle(context.Context, Command) error
}
