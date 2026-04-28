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


type LoginRequest struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Product  struct {
	Product_name string  `json:"product_name" binding:"required"`
	Product_price float64  `json:"product_price" binding:"required"`
	Product_Description string `json:"product_description" binding:"required"`	
}


func getdata(c *gin.Context) {
	user := []string{"Arbaaz", "Rahul", "Amit"}
	c.JSON(200, user)
}

func getnumbers(c *gin.Context) {
	data := []int{1, 2, 3, 4, 5}
	c.JSON(200, data)
}

func getallusers(c *gin.Context) {
	data := []map[string]string{
		{
			"name": "Arbaaz",
			"age":  "21",
		},
		{
			"name": "Rahul",
			"age":  "22",
		},
		{
			"name": "Amit",
			"age":  "23",
		},
	}
	c.JSON(200, data)
}


var users []User
var products []Product

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


func Login(c *gin.Context) {
	var loginUser LoginRequest

	if err := c.ShouldBindJSON(&loginUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data!"})
		return
	}

	for _, user := range users {
		if user.Email == loginUser.Email && bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)) == nil {
			c.JSON(http.StatusOK, gin.H{
				"message": "Login successful!",
				"user":    user.Username,
			})
			return
		}	
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
}




func AddProduct(c *gin.Context) {
var newProduct Product

if err := c.ShouldBindJSON(&newProduct); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data!"})
	return
}


products = append(products, newProduct)

c.JSON(http.StatusCreated, gin.H{
	"message": "Product added successfully!",
	"product": newProduct.Product_name,
})
}


func GetProduct(c *gin.Context) {
	if len(products) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No products found"})
		return
	}
	c.JSON(http.StatusOK, products)	
}


func UserInfo(c *gin.Context) {
var newUserInfo User	

if err :=  c.ShouldBindJSON(&newUserInfo); err != nil {
	c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid data!"})
	return	
}

for _, user := range users {
	if user.Email == newUserInfo.Email {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User email already exists"})
		return
	}
	if user.Username == newUserInfo.Username {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User username already exists"})
		return
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUserInfo.Password), 10)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	newUserInfo.Password = string(hashedPassword)
	
	users = append(users, newUserInfo)
	c.JSON(http.StatusOK, gin.H{"message": "User info added successfully"})

}

}
