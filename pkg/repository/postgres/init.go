package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"

	"github.com/kuzmrom7/feedback-service/pkg/config"
)

var db *gorm.DB

func Connect(cfg config.Database) error {
	var err error

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s port=%s  host=%s", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode, cfg.DBPort, cfg.DBHost)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	log.Println("Database connected")
	return nil
}

func Migrate() error {
	if err := db.AutoMigrate(&Review{}); err != nil {
		return err
	}

	log.Println("Migration success")
	return nil
}

func Close() {
	sqlDB, err := db.DB()
	err = sqlDB.Close()
	if err != nil {
		log.Println(err)
	}
}
