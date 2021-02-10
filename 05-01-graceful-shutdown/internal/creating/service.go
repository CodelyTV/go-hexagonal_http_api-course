package creating

import (
	"context"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/05-01-graceful-shutdown/internal"
)

// CourseService is the default CourseService interface
// implementation returned by creating.NewCourseService.
type CourseService struct {
	courseRepository mooc.CourseRepository
}

// NewCourseService returns the default Service interface implementation.
func NewCourseService(courseRepository mooc.CourseRepository) CourseService {
	return CourseService{
		courseRepository: courseRepository,
	}
}

// CreateCourse implements the creating.CourseService interface.
func (s CourseService) CreateCourse(ctx context.Context, id, name, duration string) error {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil {
		return err
	}
	return s.courseRepository.Save(ctx, course)
}
