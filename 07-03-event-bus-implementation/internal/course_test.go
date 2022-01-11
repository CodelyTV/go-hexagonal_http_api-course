package mooc

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCourse_PullEvents(t *testing.T) {
	courseID, courseName, courseDuration := "37a0f027-15e6-47cc-a5d2-64183281087e", "Test Course", "10 months"
	course, err := NewCourse(courseID, courseName, courseDuration)
	require.NoError(t, err)

	assert.GreaterOrEqual(t, len(course.PullEvents()), 1)
	assert.Len(t, course.PullEvents(), 0)
}
