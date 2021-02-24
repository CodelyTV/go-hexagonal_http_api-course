package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal/creating"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal/increasing"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal/platform/bus/inmemory"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal/platform/server"
	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
)

func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)

	courseRepository := mysql.NewCourseRepository(db, cfg.DbTimeout)

	creatingCourseService := creating.NewCourseService(courseRepository, eventBus)
	increasingCourseService := increasing.NewCourseCounterService()

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	eventBus.Subscribe(
		mooc.CourseCreatedEventType,
		creating.NewIncreaseCoursesCounterOnCourseCreated(increasingCourseService),
	)

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"localhost"`
	Port            uint          `default:"8080"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"codely"`
	DbPass    string        `default:"codely"`
	DbHost    string        `default:"localhost"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"codely"`
	DbTimeout time.Duration `default:"5s"`
}
