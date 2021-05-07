package courses

import (
	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal/fetching"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/kit/bus"
	"github.com/gin-gonic/gin"
	"net/http"
)

type getResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

// GetHandler returns an HTTP handler for courses.
func GetHandler(queryBus bus.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var queryResponse, err = queryBus.DispatchQuery(ctx, fetching.NewFetchCourseQuery())

		if err != nil {
			// Si quiero devolver error en ves de la lista se rompe me genera un error de unmarshal
			ctx.JSON(http.StatusInternalServerError, []getResponse{})
			return
		}
		courses, ok := queryResponse.([]mooc.Course)
		if ok {
			var response = make([]getResponse, 0, len(courses))
			for _, course := range courses {
				response = append(response, getResponse{
					Id:       course.ID().String(),
					Name:     course.Name().String(),
					Duration: course.Duration().String(),
				})
			}
			ctx.JSON(http.StatusOK, response)
		}
		ctx.JSON(http.StatusInternalServerError, []getResponse{})
	}
}
