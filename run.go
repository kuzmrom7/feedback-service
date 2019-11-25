package main

import (
	p "feedback-service/parser"
	"feedback-service/storage"
	"log"
)

func main() {
	if err := storage.Connect(); err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	p.Run()
}
