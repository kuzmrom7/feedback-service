package config

import (
	v "github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   Server
	Database Database
}

type Server struct {
	Port int
}

type Database struct {
	DBName     string
	DBPort     string
	DBHost     string
	DBUser     string
	DBPassword string
	DBSSLMode  string
}

func Init() (*Config, error) {
	v.AddConfigPath("contrib")
	v.SetConfigName("config")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config

	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	v.SetDefault("database.dbname", "test_db")

	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &cfg, nil
}
