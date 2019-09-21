// handlers.article.go

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	articles := GetAllArticle()

	//// Call the HTML method of the Context to render a template
	//c.HTML(
	//	// Set the HTTP status to 200 (OK)
	//	http.StatusOK,
	//	// Use the index.html template
	//	"index.html",
	//	// Pass the data that the page uses
	//	gin.H{
	//		"title":   "Home Page",
	//		"payload": articles,
	//	},
	//)
	render(c,gin.H{
		"title":"Home Page",
		"payload": articles,
	},
	"index.html")

}

func getArticle(c *gin.Context){
	//Check if article id is valid
	if articleID, err :=strconv.Atoi(c.Param("articleID"));err==nil{
		//Check id article exist
		if article,err :=GetArticleByID(articleID);err==nil{
			c.HTML(
				//Set the http status to 200 OK
				http.StatusOK,
				//Use the html template
				"article.html",
				//Pass the data to the page
				gin.H{
					"Title":article.Title,
					"payload":article,
				},
				)
		}else {
			//if article is not found abort with error
			c.AbortWithError(http.StatusNotFound,err)
		}
	}else {
		//If an invalid article ID is specified in the URL,abort with an error
		c.AbortWithStatus(http.StatusNotFound)
	}
}

func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}