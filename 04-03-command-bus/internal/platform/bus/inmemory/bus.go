package inmemory

import (
	"context"
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/kit/bus"
)

// Bus is an in-memory implementation of the bus.Bus.
type Bus struct {
	commandsHandlers map[bus.Type]bus.CommandHandler
	queryHandlers map[bus.Type]bus.QueryHandler
	eventHandlers map[bus.Type]bus.EventHandler
}

// NewCommandBus initializes a new instance of Bus.
func NewCommandBus() *Bus {
	return &Bus{
		commandsHandlers: make(map[bus.Type]bus.CommandHandler),
		queryHandlers: make(map[bus.Type]bus.QueryHandler),
		eventHandlers: make(map[bus.Type]bus.EventHandler),
	}
}

// DispatchCommand implements the bus.Bus interface.
func (b *Bus) DispatchCommand(ctx context.Context, cmd bus.Command) error {
	handler, ok := b.commandsHandlers[cmd.Type()]
	if !ok {
		return nil
	}

	// Si dejo este asincronismo sucede que el endpoint del controller responde 201, cuando en verdad fallo. No tiene más sentido que el asincrnismo lo maneje el que lo llama?
	/*go func() {
		err := handler.Handle(ctx, cmd)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
		}

	}()*/

	return handler.Handle(ctx, cmd)
}

// RegisterCommandHandler implements the bus.Bus interface.
func (b *Bus) RegisterCommandHandler(cmdType bus.Type, handler bus.CommandHandler) {
	b.commandsHandlers[cmdType] = handler
}

// DispatchQuery implements the bus.Bus interface.
func (b *Bus) DispatchQuery(ctx context.Context, query bus.Query) (bus.QueryResponse, error) {
	handler, ok := b.queryHandlers[query.Type()]
	if !ok {
		return nil, nil
	}
	// Como este es sincronico, decide el que lo usa si lo espera o no.
	return handler.Handle(ctx, query)
}

// RegisterQueryHandler implements the bus.Bus interface.
func (b *Bus) RegisterQueryHandler(queryType bus.Type, handler bus.QueryHandler) {
	b.queryHandlers[queryType] = handler
}

// DispatchEvent implements the bus.Bus interface.
func (b *Bus) DispatchEvent(ctx context.Context, event bus.Event) error {
	handler, ok := b.eventHandlers[event.Type()]
	if !ok {
		return nil
	}

	go func() {
		err := handler.Handle(ctx, event)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", event.Type(), err)
		}

	}()

	return nil
}

// RegisterEventHandler implements the bus.Bus interface.
func (b *Bus) RegisterEventHandler(cmdType bus.Type, handler bus.EventHandler) {
	b.eventHandlers[cmdType] = handler
}