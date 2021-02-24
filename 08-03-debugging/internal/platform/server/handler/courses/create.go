package courses

import (
	"errors"
	"net/http"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal/creating"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/kit/command"
	"github.com/gin-gonic/gin"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler(commandBus command.Bus) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		err := commandBus.Dispatch(ctx, creating.NewCourseCommand(
			req.ID,
			req.Name,
			req.Duration,
		))

		if err != nil {
			switch {
			case errors.Is(err, mooc.ErrInvalidCourseID),
				errors.Is(err, mooc.ErrEmptyCourseName), errors.Is(err, mooc.ErrInvalidCourseID):
				ctx.JSON(http.StatusBadRequest, err.Error())
				return
			default:
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
		}

		ctx.Status(http.StatusCreated)
	}
}
