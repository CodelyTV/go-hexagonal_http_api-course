package main

import (
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/01-03-architectured-healthcheck/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
