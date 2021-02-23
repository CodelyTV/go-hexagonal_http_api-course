package main

import (
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/07-01-publishing-domain-events/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
