package courses

import (
	"net/http"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal"
	"github.com/gin-gonic/gin"
)

type CourseResponse struct {
	Id       string
	Name     string
	Duration string
}

// CreateHandler returns an HTTP handler for courses creation.
func SearchAllHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		courses, err := courseRepository.SearchAll(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		var coursesResponse []CourseResponse

		for _, course := range courses {
			coursesResponse = append(coursesResponse, CourseResponse{
				Id:       course.ID().String(),
				Name:     course.Name().String(),
				Duration: course.Duration().String(),
			})
		}

		ctx.JSON(http.StatusOK, coursesResponse)
	}
}
