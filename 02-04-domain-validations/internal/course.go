package mooc

import (
	"context"
	"errors"
	"fmt"
	"log"

	"strings"

	"github.com/google/uuid"
)

var ErrInvalidCourseID = errors.New("invalid Course ID")
var ErrBasicConditionsCourseName = errors.New("invalid Course Name")

// CourseID represents the course unique identifier.
type CourseID struct {
	value string
}

// NewCourseID instantiate the VO for CourseID
func NewCourseID(value string) (CourseID, error) {
	v, err := uuid.Parse(value)
	if err != nil {
		return CourseID{}, fmt.Errorf("%w: %s", ErrInvalidCourseID, value)
	}

	return CourseID{
		value: v.String(),
	}, nil
}

// String type converts the CourseID into string.
func (id CourseID) String() string {
	return id.value
}

var ErrEmptyCourseName = errors.New("the field Course Name can not be empty")

// CourseName represents the course name.
type CourseName struct {
	value string
}

// NewCourseName instantiate VO for CourseName
func NewCourseName(value string) (CourseName, error) {
	if value == "" {
		return CourseName{}, ErrEmptyCourseName
	}

	// debe tener al menos un espacio en blanco
	if !strings.Contains(value, " ") {
		log.Println("no tiene espacio en blanco")
		return CourseName{}, fmt.Errorf("%w: %s", ErrBasicConditionsCourseName, value)
	}

	// su longitud minima debe ser de 10 caracteres
	var length = len([]rune(value))

	if length < 10 {
		log.Println("la longitud es de menos de 10 caracteres")
		return CourseName{}, fmt.Errorf("%w: %s", ErrBasicConditionsCourseName, value)
	}

	return CourseName{
		value: value,
	}, nil
}

// String type converts the CourseName into string.
func (name CourseName) String() string {
	return name.value
}

var ErrEmptyDuration = errors.New("the field Duration can not be empty")

// CourseDuration represents the course duration.
type CourseDuration struct {
	value string
}

func NewCourseDuration(value string) (CourseDuration, error) {
	if value == "" {
		return CourseDuration{}, ErrEmptyDuration
	}

	return CourseDuration{
		value: value,
	}, nil
}

// String type converts the CourseDuration into string.
func (duration CourseDuration) String() string {
	return duration.value
}

// CourseRepository defines the expected behaviour from a course storage.
type CourseRepository interface {
	Save(ctx context.Context, course Course) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=platform/storage/storagemocks --name=CourseRepository

// Course is the data structure that represents a course.
type Course struct {
	id       CourseID
	name     CourseName
	duration CourseDuration
}

// NewCourse creates a new course.
func NewCourse(id, name, duration string) (Course, error) {
	idVO, err := NewCourseID(id)
	if err != nil {
		return Course{}, err
	}

	nameVO, err := NewCourseName(name)
	if err != nil {
		return Course{}, err
	}

	durationVO, err := NewCourseDuration(duration)
	if err != nil {
		return Course{}, err
	}

	return Course{
		id:       idVO,
		name:     nameVO,
		duration: durationVO,
	}, nil
}

// ID returns the course unique identifier.
func (c Course) ID() CourseID {
	return c.id
}

// Name returns the course name.
func (c Course) Name() CourseName {
	return c.name
}

// Duration returns the course duration.
func (c Course) Duration() CourseDuration {
	return c.duration
}
