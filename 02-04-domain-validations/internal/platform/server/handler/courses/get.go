package courses

import (
	"net/http"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal"
	"github.com/gin-gonic/gin"
)

type getResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Duration string `json:"duration"`
}

// GetHandler returns an HTTP handler for courses.
func GetHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var courses, err = courseRepository.GetAll(ctx)
		if err != nil {
			// Si quiero devolver error en ves de la lista se rompe me genera un error de unmarshal
			ctx.JSON(http.StatusInternalServerError, []getResponse{})
			return
		}
		response := make([]getResponse, 0, len(courses))
		for _, course := range courses {
			response = append(response, getResponse{
				Id:       course.ID().String(),
				Name:     course.Name().String(),
				Duration: course.Duration().String(),
			})
		}
		ctx.JSON(http.StatusOK, response)
	}
}
