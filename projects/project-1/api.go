package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

var users []User

func getdata(c *gin.Context) {
	user := []string{"Arbaaz", "Rahul", "Amit"}
	c.JSON(200, user)
}

func getnumbers(c *gin.Context) {
	data := []int{1, 2, 3, 4, 5}
	c.JSON(200, data)
}

func getallusers(c * gin.Context){
	data := []map[string]string{
		{
			"name" : "Arbaaz",
			"age" : "21",
		},
		{
			"name" : "Rahul",
			"age" : "22",
		},
		{
			"name" : "Amit",
			"age" : "23",
		},
	}
	c.JSON(200, data)	
}

func Register(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data!"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUser.Password = string(hashedPassword)

	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully! 🎉",
		"user":    newUser.Username,
	})
}

