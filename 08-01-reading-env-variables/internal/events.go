package mooc

import (
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-01-reading-env-variables/kit/event"
)

const CourseCreatedEventType event.Type = "events.course.created"

type CourseCreatedEvent struct {
	event.BaseEvent
	id       string
	name     string
	duration string
}

func NewCourseCreatedEvent(id, name, duration string) CourseCreatedEvent {
	return CourseCreatedEvent{
		id:       id,
		name:     name,
		duration: duration,

		BaseEvent: event.NewBaseEvent(id),
	}
}

func (e CourseCreatedEvent) Type() event.Type {
	return CourseCreatedEventType
}

func (e CourseCreatedEvent) CourseID() string {
	return e.id
}

func (e CourseCreatedEvent) CourseName() string {
	return e.name
}

func (e CourseCreatedEvent) CourseDuration() string {
	return e.duration
}
