package main

// package project1

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server running 🚀",
		})
	})

	// r.GET("/users", func(c *gin.Context) {
	// 	users := []string{"Arbaaz", "Rahul", "Amit"}
	// 	c.JSON(200, users)
	// })

	r.GET("/users", getdata)
	r.GET("/getnumbers", getnumbers)
	r.GET("/getallusers", getallusers)
	r.POST("/register", Register)

	r.Run(":4000")
}
