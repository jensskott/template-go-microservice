package main

import (
	"log"

	"github.com/getndazn/template-go-microservice/router"
)

func main() {
	r := router.Setup()

	if err := r.Run(":3000"); err != nil {
		log.Fatal("cant run app")
	}
}
