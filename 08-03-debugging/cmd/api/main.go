package main

import (
	"log"

	"github.com/CodelyTV/go-hexagonal_http_api-course/08-03-debugging/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
