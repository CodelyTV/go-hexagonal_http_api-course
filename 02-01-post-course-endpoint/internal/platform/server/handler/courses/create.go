package courses

import (
	"database/sql"
	"fmt"
	"net/http"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/02-01-post-course-endpoint/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/02-01-post-course-endpoint/internal/platform/storage/mysql"
	"github.com/gin-gonic/gin"
)

const (
	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

type createRequest struct {
	ID       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

// CreateHandler returns an HTTP handler for courses creation.
func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		course := mooc.NewCourse(req.ID, req.Name, req.Duration)

		mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
		db, err := sql.Open("mysql", mysqlURI)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		courseRepository := mysql.NewCourseRepository(db)

		if err := courseRepository.Save(ctx, course); err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}
