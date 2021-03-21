package server

import (
	"fmt"
	"log"

	mooc "github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal"
	"github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal/platform/server/handler/courses"
	"github.com/CodelyTV/go-hexagonal_http_api-course/02-04-domain-validations/internal/platform/server/handler/health"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	courseRepository mooc.CourseRepository
}

func New(host string, port uint, courseRepository mooc.CourseRepository) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		courseRepository: courseRepository,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/courses", courses.CreateHandler(s.courseRepository))
	s.engine.GET("/courses", courses.GetHandler(s.courseRepository))
}
