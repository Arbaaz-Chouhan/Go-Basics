package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Ek global variable banayenge taaki pure project me DB use kar sakein
var DB *gorm.DB

func ConnectDatabase() {
	// 1. .env file ko load karna
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 2. .env file se variables read karna
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	// 3. PostgreSQL ka connection string (DSN) banana
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		dbHost, dbUser, dbPassword, dbName, dbPort)

	// 4. GORM ke zariye DB se connect karna
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database! ❌", err)
	}

	fmt.Println("Database connected successfully! ✅")

	DB = database // Global variable me assign kar diya
}
