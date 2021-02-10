package bootstrap

import (
	"database/sql"
	"fmt"

	"github.com/CodelyTV/go-hexagonal_http_api-course/04-01-application-service/internal/creating"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-01-application-service/internal/platform/server"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-01-application-service/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "codely"
	dbPass = "codely"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "codely"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	creatingCourseService := creating.NewCourseService(courseRepository)

	srv := server.New(host, port, creatingCourseService)
	return srv.Run()
}
