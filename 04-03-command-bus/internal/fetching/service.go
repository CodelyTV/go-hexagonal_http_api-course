package fetching

import (
	"context"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal"
)

// CourseService is the default CourseService interface
// implementation returned by fetching.NewCourseFetchingService.
type FetchingCourseService struct {
	courseRepository mooc.CourseRepository
}

// NewCourseService returns the default Service interface implementation.
func NewCourseFetchingService(courseRepository mooc.CourseRepository) FetchingCourseService {
	return FetchingCourseService{
		courseRepository: courseRepository,
	}
}

// CreateCourse implements the creating.CourseService interface.
func (s FetchingCourseService) GetAll(ctx context.Context) ([]mooc.Course, error) {
	return s.courseRepository.GetAll(ctx)
}
