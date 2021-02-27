package courses

import (
	"encoding/json"
	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_SearchAll(t *testing.T) {

	gin.SetMode(gin.TestMode)
	r := gin.New()

	t.Run("it returns an empty array of CourseResponse when there are no courses", func(t *testing.T) {
		var emptyCourses []mooc.Course

		courseRepository := new(storagemocks.CourseRepository)
		courseRepository.On("SearchAll", mock.Anything).Return(emptyCourses, nil)
		gin.SetMode(gin.TestMode)
		r := gin.New()
		r.GET("/courses", SearchAllHandler(courseRepository))

		req, err := http.NewRequest(http.MethodGet, "/courses", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response []CourseResponse
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			log.Fatalln(err)
		}

		assert.Equal(t, []CourseResponse(nil), response)
	})

	t.Run("it returns an array of CourseResponse when there are courses", func(t *testing.T) {

		course, _ := mooc.NewCourse("8a1c5cdc-ba57-445a-994d-aa412d23723f", "New Course", "10 months")
		var courses = []mooc.Course{course}

		courseRepository := new(storagemocks.CourseRepository)
		courseRepository.On("SearchAll", mock.Anything).Return(courses, nil)
		r.GET("/courses", SearchAllHandler(courseRepository))

		req, err := http.NewRequest(http.MethodGet, "/courses", nil)
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)

		var response []CourseResponse
		if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
			log.Fatalln(err)
		}

		var expected = convertToCourseResponse(courses)
		assert.Equal(t, expected, response)
	})
}

func convertToCourseResponse(courses []mooc.Course) []CourseResponse {
	var response []CourseResponse

	if len(courses) == 0 {
		return response
	}

	for _, course := range courses {
		response = append(response, CourseResponse{
			Id:       course.ID().String(),
			Name:     course.Name().String(),
			Duration: course.Duration().String(),
		})
	}

	return response
}
