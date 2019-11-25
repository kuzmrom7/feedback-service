package main

import (
	"feedback-service/api"
	"feedback-service/storage"
	"log"
)

func main() {
	if err := storage.Connect(); err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	err := api.Run()
	if err != nil {
		log.Fatal(err)
	}

	//p.Run()
}
