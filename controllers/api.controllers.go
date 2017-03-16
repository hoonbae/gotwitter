package controllers

import (
	. "gotwitter/core"
	. "gotwitter/models"

	"github.com/gin-gonic/gin"
)

func ApiController(e *gin.Engine) {
	router := e.Group("/api")

	router.GET("", ApiTest)
	router.GET("/users", ApiGetUsers)
}

func ApiTest(c *gin.Context) {
	c.JSON(200, struct {
		Status  int
		Message string
	}{
		200,
		"It's working!",
	})
}

func ApiGetUsers(c *gin.Context) {
	// Create a variable to store the results
	var tweet []User

	// Call the database to find all the users
	Storage.DB.Find(&tweet)

	// Return all the users as JSON
	c.JSON(200, tweet)
}
