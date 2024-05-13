package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static", "./static")

	// Define a route handler for the root path
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, You created a Web App!"})
	})

	// Define a route handler for "/about"
	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "About Us"})
	})

	// Define a route handler for "/contact"
	router.GET("/contact", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Contact Us"})
	})

	router.GET("/hello/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(http.StatusOK, gin.H{"message": "Hello, " + name + "!"})
	})

	router.GET("/profile/:username", func(c *gin.Context) {
		username := c.Param("username")
		c.HTML(http.StatusOK, "profile.html", gin.H{"username": username})
	})

	router.Run(":8080")
}
