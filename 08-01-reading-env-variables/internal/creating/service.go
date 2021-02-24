package creating

import (
	"context"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/08-01-reading-env-variables/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-01-reading-env-variables/kit/event"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService.
type CourseService struct {
	courseRepository mooc.CourseRepository
	eventBus         event.Bus
}

// NewCourseService returns the default Service interface implementation.
func NewCourseService(courseRepository mooc.CourseRepository, eventBus event.Bus) CourseService {
	return CourseService{
		courseRepository: courseRepository,
		eventBus:         eventBus,
	}
}

// CreateCourse implements the creating.CourseService interface.
func (s CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}

	if err := s.courseRepository.Save(ctx, course); err != nil {
		return err
	}

	return s.eventBus.Publish(ctx, course.PullEvents())
}
