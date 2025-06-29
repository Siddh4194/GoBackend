package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// InitDB initializes the DB only once (singleton)
func InitDB() *gorm.DB {
	once.Do(func() {
		fmt.Println("🔗 Connecting to PostgreSQL...")
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASS"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)

		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("❌ Failed to connect to database:", err)
		}
		log.Println("✅ Connected to PostgreSQL")
	})
	return db
}

// GetDB safely returns the singleton DB instance
func GetDB() *gorm.DB {
	if db == nil {
		log.Println("DB not initialized. Calling InitDB().")
        return InitDB()
	}
	return db
}
