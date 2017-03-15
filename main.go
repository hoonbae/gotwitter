package main

import (
	"github.com/gin-gonic/gin"
	"twitter/controllers"

	. "twitter/core"
)

func LoadCore() {
	Storage.Start()
}

func LoadControllers(e *gin.Engine) {
	controllers.HomeController(e)
}

func main() {
	e := gin.Default()

	LoadCore()

	e.LoadHTMLGlob("templates/**/*")
	LoadControllers(e)
	e.Run(":8080")
}
