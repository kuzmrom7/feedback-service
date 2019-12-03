package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

var db *sqlx.DB

func Connect() error {
	var err error

	db, err = sqlx.Connect("postgres", "user=root password=arbuz dbname=feedback_service sslmode=disable")
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
