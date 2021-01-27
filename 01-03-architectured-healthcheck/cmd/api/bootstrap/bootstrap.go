package bootstrap

import "github.com/CodelyTV/go-hexagonal_http_api-course/01-03-architectured-healthcheck/internal/platform/server"

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
