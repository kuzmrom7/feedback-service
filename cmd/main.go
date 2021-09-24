package main

import (
	"fmt"
	"github.com/kuzmrom7/feedback-service/pkg/api"
	"github.com/kuzmrom7/feedback-service/pkg/parser"
	"github.com/kuzmrom7/feedback-service/pkg/storage"
	"log"

	"github.com/kuzmrom7/feedback-service/pkg/config"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Database)

	if err := storage.Connect(cfg.Database); err != nil {
		log.Fatal(err)
	}
	if err = storage.Migrate(); err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	go func() {
		parser.Run()
	}()

	if err = api.Run(cfg.Server); err != nil {
		log.Fatal(err)
	}
}
