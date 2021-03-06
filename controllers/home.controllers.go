package controllers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"

	. "gotwitter/core"
	. "gotwitter/models"
)

func HomeController(e *gin.Engine) {
	router := e.Group("/home")

	router.GET("", HomeIndex)
	router.POST("", SubmitTweet)
}

func HomeIndex(c *gin.Context) {
	c.HTML(200, "home/index.html", pongo2.Context{})
}

func SubmitTweet(c *gin.Context) {
	tweet := c.PostForm("tweet")

	newTweet := Tweet{
		TweetMessage: tweet,
	}

	Storage.DB.Create(&newTweet)

	c.Redirect(302, "/home")
}
