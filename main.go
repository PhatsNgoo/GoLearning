package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
var router *gin.Engine


func main()  {
	//Creates a default route
	router :=gin.Default()

	//Load templates directory
	router.LoadHTMLGlob("templates/*")

	//Initial routes

	// Handle the index route
	router.GET("/", func(c *gin.Context) {
		articles := GetAllArticle()

		// Call the HTML method of the Context to render a template
		c.HTML(
			// Set the HTTP status to 200 (OK)
			http.StatusOK,
			// Use the index.html template
			"index.html",
			// Pass the data that the page uses
			gin.H{
				"title":   "Home Page",
				"payload": articles,
			},
		)

	})

	//Start serving the application
	router.Run()

}

