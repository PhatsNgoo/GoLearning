package main

import (
	"github.com/gin-gonic/gin"
)
var router *gin.Engine


func main()  {
	//Creates a default route
	router =gin.Default()

	//Load templates directory
	router.LoadHTMLGlob("templates/*")

	//Initial routes
	go initializeRoutes()
	// Handle the index route


	//Start serving the application
	router.Run()

}

