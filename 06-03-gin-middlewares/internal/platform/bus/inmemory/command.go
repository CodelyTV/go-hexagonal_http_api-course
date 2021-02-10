package inmemory

import (
	"context"
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/06-03-gin-middlewares/kit/command"
)

// CommandBus is an in-memory implementation of the command.Bus.
type CommandBus struct {
	handlers map[command.Type]command.Handler
}

// NewCommandBus initializes a new instance of CommandBus.
func NewCommandBus() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

// Dispatch implements the command.Bus interface.
func (b *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := b.handlers[cmd.Type()]
	if !ok {
		return nil
	}

	go func() {
		err := handler.Handle(ctx, cmd)
		if err != nil {
			log.Printf("Error while handling %s - %s\n", cmd.Type(), err)
		}

	}()

	return nil
}

// Register implements the command.Bus interface.
func (b *CommandBus) Register(cmdType command.Type, handler command.Handler) {
	b.handlers[cmdType] = handler
}
