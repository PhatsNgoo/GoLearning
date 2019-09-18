package main

import "net/http"
import "github.com/gin-gonic/gin"

var router *gin.Engine

func main()  {
	//Creates a default route
	router :=gin.Default()

	//Load templates directory
	router.LoadHTMLGlob("templates/*")

	//Define router handler
	router.GET("/", func(c *gin.Context) {
		//Call the html method
		c.HTML(
			//Set the HTTP status to 200 OK
			http.StatusOK,
			//Use Index html file
			"index.html",
			//Pass the data that the page uses(in this case 'title')
			gin.H{
				"title":"HomePage",
			},
			)
	})
	//Start serving the application
	router.Run()

}
