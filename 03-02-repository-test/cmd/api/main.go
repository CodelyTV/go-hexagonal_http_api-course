package main

import (
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/03-02-repository-test/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
