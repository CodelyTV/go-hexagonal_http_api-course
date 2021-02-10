package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/06-03-gin-middlewares/internal"
	"github.com/huandu/go-sqlbuilder"
)

// CourseRepository is a MySQL mooc.CourseRepository implementation.
type CourseRepository struct {
	db        *sql.DB
	dbTimeout time.Duration
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB, dbTimeout time.Duration) *CourseRepository {
	return &CourseRepository{
		db:        db,
		dbTimeout: dbTimeout,
	}
}

// Save implements the mooc.CourseRepository interface.
func (r *CourseRepository) Save(ctx context.Context, course mooc.Course) error {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.InsertInto(sqlCourseTable, sqlCourse{
		ID:       course.ID().String(),
		Name:     course.Name().String(),
		Duration: course.Duration().String(),
	}).Build()

	ctxTimeout, cancel := context.WithTimeout(ctx, r.dbTimeout)
	defer cancel()

	_, err := r.db.ExecContext(ctxTimeout, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}
