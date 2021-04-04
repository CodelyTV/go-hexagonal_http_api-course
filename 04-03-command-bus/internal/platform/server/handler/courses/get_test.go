package courses

import (
	"encoding/json"
	"errors"
	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal/fetching"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal/platform/bus/inmemory"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHandler(t *testing.T) {
	tests := map[string]struct {
		mockData []mooc.Course
		mockError error
		expectedResponse []getResponse
		expectedStatus int
	}{
		"Empty repo return 200 with empty list": {
			mockData: []mooc.Course{},
			mockError: nil,
			expectedStatus:http.StatusOK,
			expectedResponse: []getResponse{}},
		"Repo error return 500": {
			mockData: []mooc.Course{},
			mockError: errors.New("the field Duration can not be empty"),
			expectedStatus:http.StatusInternalServerError,
			expectedResponse: []getResponse{}},
		"Fully repo return 200 with list of courses":{
			mockData: []mooc.Course{ mockCourse("8a1c5cdc-ba57-445a-994d-aa412d23723f", "Courses Complete", "123")},
			mockError: nil,
			expectedStatus:http.StatusOK,
			expectedResponse: []getResponse{{Id: "8a1c5cdc-ba57-445a-994d-aa412d23723f", Name: "Courses Complete", Duration: "123"}}},
	}
	for key, value := range tests {
		t.Run(key, func(t *testing.T) {
			courseRepository := storagemocks.CourseRepository{}
			bus := inmemory.NewCommandBus()
			fetchingCourseService := fetching.NewCourseFetchingService(&courseRepository)
			fetchingCourseQueryHandler := fetching.NewCourseQueryHandler(fetchingCourseService)
			bus.RegisterQueryHandler(fetching.CourseQueryType, fetchingCourseQueryHandler)

			courseRepository.On(
				"GetAll",
				mock.Anything,
			).Return(value.mockData, value.mockError)
			gin.SetMode(gin.TestMode)
			r := gin.New()
			r.GET("/courses", GetHandler(bus))

			req, err := http.NewRequest(http.MethodGet, "/courses", nil)
			require.NoError(t, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, value.expectedStatus, res.StatusCode)
			var response []getResponse
			if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
				log.Fatalln(err)
			}

			assert.Equal(t, value.expectedResponse, response)
		})
	}
}

func mockCourse(id string, name string, duration string) mooc.Course {
	course, err := mooc.NewCourse(id, name, duration)
	if err != nil{
		log.Fatalln(err)
	}
	return course
}
