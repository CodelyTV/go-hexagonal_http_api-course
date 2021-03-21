package courses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal/platform/storage/storagemocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	courseRepository := new(storagemocks.CourseRepository)
	courseRepository.On("Save", mock.Anything, mock.AnythingOfType("mooc.Course")).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(courseRepository))

	tests := map[string]struct {
		id string
		name string
		duration string
		want int
	}{
		"given an invalid request it returns 400": {name: "Demo Course", duration: "10 months", want: http.StatusBadRequest},
		"given an invalid name it returns 412": {id: "8a1c5cdc-ba57-445a-994d-aa412d23723f", name: "412 Course", duration: "10 months", want: http.StatusPreconditionFailed},
		"given a valid request it returns 201": {id: "8a1c5cdc-ba57-445a-994d-aa412d23723f", name: "Demo Course", duration: "10 months", want: http.StatusCreated},
	}
	for key, value := range tests {
		t.Run(key, func(t *testing.T) {

			createCourseReq := createRequest{
				ID: value.id,
				Name:     value.name,
				Duration: value.duration,
			}

			b, err := json.Marshal(createCourseReq)
			require.NoError(t, err)

			req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
			require.NoError(t, err)

			rec := httptest.NewRecorder()
			r.ServeHTTP(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			assert.Equal(t, value.want, res.StatusCode)
		})
	}
}
