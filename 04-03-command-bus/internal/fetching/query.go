package fetching

import (
	"context"
	"errors"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/kit/bus"
)

const CourseQueryType bus.Type = "bus.fetching.courses"

// CourseQuery is the bus dispatched to create a new course.
type CourseQuery struct {
}

// NewFetchCourseQuery creates a new CourseQuery.
func NewFetchCourseQuery() CourseQuery {
	return CourseQuery{}
}

func (c CourseQuery) Type() bus.Type {
	return CourseQueryType
}

// CourseQueryHandler is the bus handler
// responsible for fetching courses.
type CourseQueryHandler struct {
	service FetchingCourseService
}

// NewCourseQueryHandler initializes a new NewCourseQueryHandler.
func NewCourseQueryHandler(service FetchingCourseService) CourseQueryHandler {
	return CourseQueryHandler{
		service: service,
	}
}

// Handle implements the bus.QueryHandler interface.
func (h CourseQueryHandler) Handle(ctx context.Context, query bus.Query) (bus.QueryResponse, error) {
	_, ok := query.(CourseQuery)
	if !ok {
		return nil, errors.New("unexpected bus")
	}

	return h.service.GetAll(ctx)
}
