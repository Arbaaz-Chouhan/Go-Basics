package main

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", CreateUser)
	r.GET("/get-all-users", GetAllUsers)

	// Product routes
	r.POST("/products", CreateNewProduct)
	r.GET("/products", GetAllProducts)
	r.DELETE("/products/:id", DeleteProductById)
}
