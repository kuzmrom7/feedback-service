package main

import (
	"log"

	v "github.com/spf13/viper"

	"feedback-service/pkg/api"
	"feedback-service/pkg/config"
	"feedback-service/pkg/parser"
	"feedback-service/pkg/storage"
)

func main() {

	v.SetConfigName("config")
	v.SetConfigType("yml")
	v.AddConfigPath("./contrib")
	v.AutomaticEnv()

	var cfg *config.Configurations

	if err := v.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)

	}
	v.SetDefault("database.dbname", "test_db")

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
	}

	if err := storage.Connect(cfg.Database); err != nil {
		log.Fatal(err)
	}
	defer storage.Close()

	parser.Run()

	err = api.Run(cfg.Server)
	if err != nil {
		log.Fatal(err)
	}

}
