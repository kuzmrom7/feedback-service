package storage

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"feedback-service/pkg/config"
)

var db *sqlx.DB

func Connect(cfg *config.DatabaseConfigurations) error {
	var err error

	src := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)

	db, err = sqlx.Connect("postgres", src)
	if err != nil {
		return err
	}

	log.Println("database connected")
	return nil
}

func Close() {
	if err := db.Close(); err != nil {
		log.Println(err)
	}
}
