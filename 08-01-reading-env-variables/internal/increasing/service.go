package increasing

type CourseCounterService struct{}

func NewCourseCounterService() CourseCounterService {
	return CourseCounterService{}
}

func (s CourseCounterService) Increase(id string) error {
	return nil
}
