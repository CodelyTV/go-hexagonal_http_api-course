package mysql

const (
	sqlCourseTable = "courses"
)

type sqlCourse struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Duration string `db:"duration"`
}
