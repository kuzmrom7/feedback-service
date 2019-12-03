package main

import (
	"log"

	"feedback-service/pkg/api"
	"feedback-service/pkg/parser"
	"feedback-service/pkg/storage"
)

func main() {
	if err := storage.Connect(); err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	parser.Run()

	err := api.Run()
	if err != nil {
		log.Fatal(err)
	}

}
