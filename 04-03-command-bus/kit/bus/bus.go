package bus

import "context"

// Bus defines the expected behaviour from a bus bus.
type Bus interface {
	// DispatchCommand is the method used to dispatch new commands.
	DispatchCommand(context.Context, Command) error
	// RegisterCommandHandler is the method used to register a new command handler.
	RegisterCommandHandler(Type, CommandHandler)
	// DispatchQuery is the method used to dispatch new queries.
	DispatchQuery(context.Context, Query) (QueryResponse, error)
	// RegisterQueryHandler is the method used to register a new query handler.
	RegisterQueryHandler(Type, QueryHandler)
	// DispatchEvent is the method used to dispatch new event.
	DispatchEvent(context.Context, Event) error
	// RegisterEventHandler is the method used to register a new event handler.
	RegisterEventHandler(Type, EventHandler)
}

//go:generate mockery --case=snake --outpkg=busmocks --output=busmocks --name=Bus

// Type represents an application bus type.
type Type string

// Command represents an application bus.
type Command interface {
	Type() Type
}

// CommandHandler defines the expected behaviour from a bus handler.
type CommandHandler interface {
	Handle(context.Context, Command) error
}

// Query represents an application bus.
type Query interface {
	Type() Type
}

type QueryResponse interface{}

// QueryHandler defines the expected behaviour from a bus handler.
type QueryHandler interface {
	Handle(context.Context, Query) (QueryResponse, error)
}


// Event represents an application bus.
type Event interface {
	Type() Type
}

// EventHandler defines the expected behaviour from a bus handler.
type EventHandler interface {
	Handle(context.Context, Event) error
}
