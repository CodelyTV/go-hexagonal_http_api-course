package mysql

import (
	"context"
	"database/sql"
	"fmt"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal"
	"github.com/huandu/go-sqlbuilder"
)

// CourseRepository is a MySQL mooc.CourseRepository implementation.
type CourseRepository struct {
	db *sql.DB
}

// NewCourseRepository initializes a MySQL-based implementation of mooc.CourseRepository.
func NewCourseRepository(db *sql.DB) *CourseRepository {
	return &CourseRepository{
		db: db,
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

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error trying to persist course on database: %v", err)
	}

	return nil
}

func (r *CourseRepository) SearchAll(ctx context.Context) ([]mooc.Course, error) {
	courseSQLStruct := sqlbuilder.NewStruct(new(sqlCourse))
	query, args := courseSQLStruct.SelectFrom(sqlCourseTable).Build()

	rows, err := r.db.QueryContext(ctx, query, args...)

	if err != nil {
		return []mooc.Course{}, fmt.Errorf("error trying to search all courses on database: %v", err)
	}

	var courses []mooc.Course
	for rows.Next() {
		var id string
		var name string
		var duration string

		if err := rows.Scan(&id, &name, &duration); err != nil {
			return []mooc.Course{}, fmt.Errorf("error trying to search all courses on database: %v", err)
		}

		course, err := mooc.NewCourse(id, name, duration)

		if err != nil {
			return []mooc.Course{}, fmt.Errorf("error trying to create new course in repository : %v", err)
		}

		courses = append(courses, course)
	}

	return courses, nil
}
