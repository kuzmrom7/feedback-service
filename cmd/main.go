package main

import (
	"fmt"
	"github.com/kuzmrom7/feedback-service/pkg/parser"
	"github.com/kuzmrom7/feedback-service/pkg/repository"
	"github.com/kuzmrom7/feedback-service/pkg/repository/postgres"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/kuzmrom7/feedback-service/pkg/config"
	"github.com/kuzmrom7/feedback-service/pkg/server"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := initDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}
	defer closeDB(db)

	reviewsRepository := postgres.NewReviewsRepository(db)
	prs := parser.NewParser(cfg.Parser,reviewsRepository)
	srv := server.New(cfg.Server, reviewsRepository)

	go func() {
		prs.Run()
	}()

	if err = srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func initDB(cfg config.Database) (*gorm.DB, error) {
	var db *gorm.DB

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s port=%s  host=%s", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode, cfg.DBPort, cfg.DBHost)
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	log.Println("Database connected")

	if err := db.AutoMigrate(&repository.Review{}); err != nil {
		log.Println("Migration failed")

		return nil, err
	}

	log.Println("Migration success")

	return db, nil
}

func closeDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	err = sqlDB.Close()
	if err != nil {
		log.Println(err)
	}
}
