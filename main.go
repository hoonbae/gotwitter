package main

import (
	"github.com/gin-gonic/gin"
	"gotwitter/controllers"

	. "gotwitter/core"

	"github.com/robvdl/pongo2gin"
)

func LoadCore() {
	Storage.Start()
}

func LoadControllers(e *gin.Engine) {
	controllers.HomeController(e)
	controllers.ApiController(e)
}

func main() {
	e := gin.Default()

	LoadCore()

	e.LoadHTMLGlob("templates/**/*")
	LoadControllers(e)

	e.Static("/static", "static")

	e.HTMLRender = pongo2gin.Default()
	e.Run(":8080")
}
