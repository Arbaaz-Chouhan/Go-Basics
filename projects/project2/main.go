package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server...")
	// Database connect function call kar rahe hain
	ConnectDatabase()
	// Create the 'users' table in the database
	DB.AutoMigrate(&User{}, &Product{}, &student{})
	fmt.Println("Table created successfully!")

	r := gin.Default()

	SetupRoutes(r)

	r.Run(":9000")
}
