package server

import (
	"fmt"
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal/platform/server/handler/courses"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/internal/platform/server/handler/health"
	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/kit/command"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	// deps
	commandBus command.Bus
}

func New(host string, port uint, commandBus command.Bus) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		commandBus: commandBus,
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
	s.engine.POST("/courses", courses.CreateHandler(s.commandBus))
}
