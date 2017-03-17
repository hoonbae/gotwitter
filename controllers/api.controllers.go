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
	router.GET("/tweets", ApiGetTweets)
	router.GET("/users/:id", ApiGetUserById)
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
	var user []User

	// Call the database to find all the users
	Storage.DB.Preload("Tweets").Find(&user)

	// Return all the users as JSON
	c.JSON(200, user)
}

func ApiGetTweets(c *gin.Context) {
	var tweet []Tweet

	Storage.DB.Preload("User").Find(&tweet)

	c.JSON(200, tweet)
}

func ApiGetUserById(c *gin.Context) {
	var user []User
	id := c.Param("id")
	Storage.DB.Preload("Tweets").Where("id=?", id).Find(&user)

	c.JSON(200, user)
}
