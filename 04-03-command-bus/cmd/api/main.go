package main

import (
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/04-03-command-bus/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
