package courses

import (
	"errors"
	"log"
	"net/http"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(courseRepository mooc.CourseRepository) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course, err := mooc.NewCourse(req.ID, req.Name, req.Duration)

		if err != nil {
			// ver aqui devolver distintos codigos de errores segun cada VO
			// CourseId   -> 406
			// CourseName -> 410
			log.Println(err.Error())

			if errors.Is(err, mooc.ErrInvalidCourseID) {
				ctx.JSON(http.StatusNotAcceptable, err.Error())
				return
			}
			if errors.Is(err, mooc.ErrBasicConditionsCourseName) {
				ctx.JSON(http.StatusGone, err.Error())
				return
			}

			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		if err := courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
