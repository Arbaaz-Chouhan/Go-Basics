package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	ID    uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"not null" json:"email"`
}

type Product struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"not null" json:"name"`
	Price       float64 `gorm:"not null" json:"price"`
	Description string  `gorm:"not null " json:"description"`
}

func CreateUser(c *gin.Context) {
	var user User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	DB.Create(&user)
	c.JSON(200, user)
}

func GetAllUsers(c *gin.Context) {
	var user []User

	DB.Find(&user)

	c.JSON(200, user)
}

func CreateNewProduct(c *gin.Context) {
	var newProduct Product

	err := c.ShouldBindJSON(&newProduct)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Create(&newProduct)
	c.JSON(http.StatusOK, newProduct)

}

func GetAllProducts(c *gin.Context) {
	var product []Product
	DB.Find(&product)
	c.JSON(http.StatusOK, product)
	return

}

func DeleteProductById(c *gin.Context) {
	var product Product
	id := c.Param("id")

	// 1. Check karein ki product exists karta hai ya nahi
	if err := DB.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// 2. Agar mil gaya, toh delete karein
	DB.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"message": "Product deleted successfully"})
}
