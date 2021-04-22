package courses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/kit/bus/busmocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandler_Create(t *testing.T) {
	bus := new(busmocks.Bus)
	bus.On(
		"DispatchCommand",
		mock.Anything,
		mock.AnythingOfType("creating.CourseCommand"),
	).Return(nil)

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/courses", CreateHandler(bus))

	t.Run("given an invalid request it returns 400", func(t *testing.T) {
		createCourseReq := createRequest{
			Name:     "Demo Course",
			Duration: "10 months",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusBadRequest, res.StatusCode)
	})

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		createCourseReq := createRequest{
			ID:       "8a1c5cdc-ba57-445a-994d-aa412d23723f",
			Name:     "Demo Course",
			Duration: "10 months",
		}

		b, err := json.Marshal(createCourseReq)
		require.NoError(t, err)

		req, err := http.NewRequest(http.MethodPost, "/courses", bytes.NewBuffer(b))
		require.NoError(t, err)

		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)

		res := rec.Result()
		defer res.Body.Close()

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})
}
