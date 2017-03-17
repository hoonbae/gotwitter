package controllers

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
)

func HomeController(e *gin.Engine) {
	router := e.Group("/home")

	router.GET("", HomeIndex)
}

func HomeIndex(c *gin.Context) {
	c.HTML(200, "home/index.html", pongo2.Context{
		"title":    "GoTwitter",
		"username": "username_placeholder",
	})
}
