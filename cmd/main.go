package main

import (
	"fmt"
	"github.com/kuzmrom7/feedback-service/pkg/parser"
	"github.com/kuzmrom7/feedback-service/pkg/repository/postgres"
	"log"

	"github.com/kuzmrom7/feedback-service/pkg/config"
	"github.com/kuzmrom7/feedback-service/pkg/server"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Database)

	if err := postgres.Connect(cfg.Database); err != nil {
		log.Fatal(err)
	}
	if err = postgres.Migrate(); err != nil {
		log.Fatal(err)
	}
	defer postgres.Close()

	go func() {
		parser.Run()
	}()

	s := server.New(cfg.Server)
	if err = s.Run(); err != nil {
		log.Fatal(err)
	}
}
