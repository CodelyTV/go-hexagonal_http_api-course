package main

import (
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/08-01-reading-env-variables/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
