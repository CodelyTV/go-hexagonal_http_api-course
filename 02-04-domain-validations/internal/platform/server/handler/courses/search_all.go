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

		var courses [1]CourseResponse

		course1, _ := mooc.NewCourse("8a1c5cdc-ba57-445a-994d-aa412d23723f", "New Course", "10 hours")
		courses[0] = CourseResponse{
			Id:       course1.ID().String(),
			Name:     course1.Name().String(),
			Duration: course1.Duration().String(),
		}

		ctx.JSON(http.StatusOK, courses)
	}
}
