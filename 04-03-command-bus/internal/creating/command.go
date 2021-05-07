package creating

import (
	"context"
	"errors"

	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/kit/bus"
)

const CourseCommandType bus.Type = "bus.creating.course"

// CourseCommand is the bus dispatched to create a new course.
type CourseCommand struct {
	id       string
	name     string
	duration string
}

// NewCourseCommand creates a new CourseCommand.
func NewCourseCommand(id, name, duration string) CourseCommand {
	return CourseCommand{
		id:       id,
		name:     name,
		duration: duration,
	}
}

func (c CourseCommand) Type() bus.Type {
	return CourseCommandType
}

// CourseCommandHandler is the bus handler
// responsible for creating courses.
type CourseCommandHandler struct {
	service CourseService
}

// NewCourseCommandHandler initializes a new CourseCommandHandler.
func NewCourseCommandHandler(service CourseService) CourseCommandHandler {
	return CourseCommandHandler{
		service: service,
	}
}

// Handle implements the bus.CommandHandler interface.
func (h CourseCommandHandler) Handle(ctx context.Context, cmd bus.Command) error {
	createCourseCmd, ok := cmd.(CourseCommand)
	if !ok {
		return errors.New("unexpected bus")
	}

	return h.service.CreateCourse(
		ctx,
		createCourseCmd.id,
		createCourseCmd.name,
		createCourseCmd.duration,
	)
}
