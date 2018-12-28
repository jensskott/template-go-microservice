package main

import (
	"log"
	"os"

	"github.com/jensskott/template-go-microservice/lib/router"
)

func main() {
	r := router.Setup()

	port := os.Getenv("PORT")

	if port == "" {
		port = ":3000"
	}

	if err := r.Run(port); err != nil {
		log.Fatal("cant run app")
	}
}
