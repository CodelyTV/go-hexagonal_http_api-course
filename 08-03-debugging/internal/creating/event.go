package creating

import (
	"context"
	"errors"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal/increasing"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/kit/event"
)

type IncreaseCoursesCounterOnCourseCreated struct {
	increasingService increasing.CourseCounterService
}

func NewIncreaseCoursesCounterOnCourseCreated(increaserService increasing.CourseCounterService) IncreaseCoursesCounterOnCourseCreated {
	return IncreaseCoursesCounterOnCourseCreated{
		increasingService: increaserService,
	}
}

func (e IncreaseCoursesCounterOnCourseCreated) Handle(_ context.Context, evt event.Event) error {
	courseCreatedEvt, ok := evt.(mooc.CourseCreatedEvent)
	if !ok {
		return errors.New("unexpected event")
	}

	return e.increasingService.Increase(courseCreatedEvt.ID())
}
