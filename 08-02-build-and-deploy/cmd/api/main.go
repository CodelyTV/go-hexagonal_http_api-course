package main

import (
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/08-02-build-and-deploy/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
